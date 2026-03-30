package examplee2ekind_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
	"sigs.k8s.io/e2e-framework/pkg/features"
	"sigs.k8s.io/e2e-framework/support/kind"
	"sigs.k8s.io/e2e-framework/third_party/helm"
)

var (
	testenv   env.Environment
	namespace string
	version   string
)

func TestMain(m *testing.M) {
	testenv = env.New()
	kindClusterName := envconf.RandomName("e2e-test", 16)
	namespace = "e2e-kind"
	if v := os.Getenv("VERSION"); v != "" {
		version = v
	} else {
		version = "1.0.0"
	}

	// pre-test setup of kind cluster
	testenv.Setup(
		envfuncs.CreateCluster(kind.NewProvider(), kindClusterName),
		envfuncs.CreateNamespace(namespace),
		// you can have multiple LoadDockerImageToCluster to load multiple images
		envfuncs.LoadDockerImageToCluster(kindClusterName, fmt.Sprintf("example-e2e-kind:%s", version)),
	)

	// post-test teardown kind cluster
	testenv.Finish(
		envfuncs.DeleteNamespace(namespace),
		envfuncs.DestroyCluster(kindClusterName),
	)
	os.Exit(testenv.Run(m))
}

func TestExample(t *testing.T) {
	// test code here
	t.Run("Test Example", func(t *testing.T) {
		t.Logf("Running test with version: %s", version)
		f := features.New("install applications").Setup(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			helmMgr := helm.New(cfg.KubeconfigFile())
			// add more args if needed
			// like disable serviceMonitor deployment if you do not have plans to add related CRD
			err := helmMgr.RunInstall(helm.WithName("example-e2e-kind"), helm.WithNamespace(namespace), helm.WithReleaseName("example-e2e-kind"), helm.WithChart("charts/example-e2e-kind"), helm.WithArgs("--set", "image.tag="+version, "--set", "image.repository=example-e2e-kind"))
			if err != nil {
				t.Fatalf("Failed to install example-e2e-kind helm chart: %v", err)
			}
			// waiting pod be ready, can be less time.
			time.Sleep(60 * time.Second)
			return ctx
		}).Assess("checking pod", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			// test application running on kind
			podList := corev1.PodList{}
			err := cfg.Client().Resources(namespace).List(ctx, &podList)
			if err != nil {
				t.Fatalf("Error listing pods: %v", err)
			}
			podName := ""
			for _, pod := range podList.Items {
				if strings.HasPrefix(pod.Name, "example-e2e-kind") {
					t.Logf("Found pod %v with status %v", pod.Name, pod.Status.Phase)
					podName = pod.Name
					break
				}
			}
			if podName == "" {
				t.Fatalf("Failed to find pod with name prefix example-e2e-kind")
			}
			// use podName here to run commands if wanted
			t.Logf("Checking pod %v", podName)

			return ctx
		}).Teardown(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			helmMgr := helm.New(cfg.KubeconfigFile())
			err := helmMgr.RunUninstall(helm.WithName("example-e2e-kind"), helm.WithNamespace(namespace))
			if err != nil {
				t.Fatalf("Failed to uninstall example-e2e-kind helm chart: %v", err)
			}
			return ctx
		}).Feature()
		testenv.Test(t, f)
	})

}
