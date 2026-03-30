# Troubleshoot 

## go mod tidy

Running go mod tidy and recevind this error.

```bash
$ go mod tidy
go: finding module for package k8s.io/api/core/v1
go: finding module for package sigs.k8s.io/e2e-framework/pkg/envconf
go: finding module for package sigs.k8s.io/e2e-framework/pkg/features
go: finding module for package sigs.k8s.io/e2e-framework/pkg/envfuncs
go: finding module for package sigs.k8s.io/e2e-framework/support/kind
go: finding module for package sigs.k8s.io/e2e-framework/pkg/env
go: finding module for package sigs.k8s.io/e2e-framework/third_party/helm
go: found k8s.io/api/core/v1 in k8s.io/api v0.35.3
go: found sigs.k8s.io/e2e-framework/pkg/env in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/envconf in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/envfuncs in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/features in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/support/kind in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/third_party/helm in sigs.k8s.io/e2e-framework v0.6.0
go: finding module for package github.com/Masterminds/semver/v3
go: finding module for package k8s.io/api/storagemigration/v1alpha1
go: finding module for package k8s.io/api/networking/v1alpha1
go: found github.com/Masterminds/semver/v3 in github.com/Masterminds/semver/v3 v3.4.0
go: finding module for package k8s.io/api/storagemigration/v1alpha1
go: finding module for package k8s.io/api/networking/v1alpha1
go: example-e2e-kind tested by
	example-e2e-kind.test imports
	sigs.k8s.io/e2e-framework/pkg/envfuncs imports
	sigs.k8s.io/e2e-framework/klient/decoder imports
	k8s.io/client-go/kubernetes/scheme imports
	k8s.io/api/networking/v1alpha1: module k8s.io/api@latest found (v0.35.3), but does not contain package k8s.io/api/networking/v1alpha1
go: example-e2e-kind tested by
	example-e2e-kind.test imports
	sigs.k8s.io/e2e-framework/pkg/envfuncs imports
	sigs.k8s.io/e2e-framework/klient/decoder imports
	k8s.io/client-go/kubernetes/scheme imports
	k8s.io/api/storagemigration/v1alpha1: module k8s.io/api@latest found (v0.35.3), but does not contain package k8s.io/api/storagemigration/v1alpha1
```

It happened in 30.03.2026 and it can change, please check out [e2e-framework issues page](https://github.com/kubernetes-sigs/e2e-framework/issues) , and what I did to solve. I edited go.mod and added this:

```
module example-e2e-kind

go 1.26.1

require (
    k8s.io/client-go v0.35.1 
)
```

Run `tidy` again:
```bash
$ go mod tidy
go: finding module for package sigs.k8s.io/e2e-framework/pkg/features
go: finding module for package sigs.k8s.io/e2e-framework/third_party/helm
go: finding module for package github.com/kr/text
go: finding module for package sigs.k8s.io/e2e-framework/support/kind
go: finding module for package sigs.k8s.io/e2e-framework/pkg/env
go: finding module for package sigs.k8s.io/e2e-framework/pkg/envfuncs
go: finding module for package sigs.k8s.io/e2e-framework/pkg/envconf
go: found sigs.k8s.io/e2e-framework/pkg/env in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/envconf in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/envfuncs in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/pkg/features in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/support/kind in sigs.k8s.io/e2e-framework v0.6.0
go: found sigs.k8s.io/e2e-framework/third_party/helm in sigs.k8s.io/e2e-framework v0.6.0
go: found github.com/kr/text in github.com/kr/text v0.2.0
```