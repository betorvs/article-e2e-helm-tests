# Article: Stop Guessing: End-to-End Testing for Helm Charts with Kubernetes-sigs e2e-framework

This is a example repository which uses end-to-end Kubernetes Sigs Framework to test helm chart using Golang.

## Requirements

- [e2e-framework](https://github.com/kubernetes-sigs/e2e-framework/tree/main)
- [helm](https://helm.sh/docs/)
- [kind](https://kind.sigs.k8s.io/)
- [go](https://go.dev/)

## Run manually

You can test it with this versioned Helm chart, but I recommend cleaning it up and starting again with the command: `rm -rf charts go.*`. Use these previous files for reference in case of failure.

Then, start a new Helm and download Go dependencies.

```bash
go mod init example-e2e-kind
go mod tidy
mkdir -p charts
helm create charts/example-e2e-kind
```

In case of having issues with `go mod tidy`, consult [troubleshoot](./Troubleshoot.md)

Run the test:
```bash
docker build -t example-e2e-kind:test1 -f Dockerfile .
VERSION=test1 go test -count=1 -timeout=10m -v ./...
```

Output example:
```bash
$ VERSION=test1 go test -count=1 -timeout=10m -v ./...
=== RUN   TestExample
=== RUN   TestExample/Test_Example
    example_kind_test.go:55: Running test with version: test1
=== RUN   TestExample/Test_Example/install_applications
=== RUN   TestExample/Test_Example/install_applications/checking_pod
    example_kind_test.go:77: Found pod example-e2e-kind-6cd8f6977f-87njr with status Running
    example_kind_test.go:86: Checking pod example-e2e-kind-6cd8f6977f-87njr
--- PASS: TestExample (60.52s)
    --- PASS: TestExample/Test_Example (60.52s)
        --- PASS: TestExample/Test_Example/install_applications (60.52s)
            --- PASS: TestExample/Test_Example/install_applications/checking_pod (0.02s)
PASS
ok  	example-e2e-kind	88.262s
?   	example-e2e-kind/cmd/example-e2e-kind	[no test files]
```

## Run using task

It's require [task](https://taskfile.dev) installed.

```bash
$ task init
$ task tidy
$ task e2e-tests
```

## Tips

- Change helm chart to use correct deployment port, health check configuration.
- Remove `cmd/example` application and adds your own. Don't forgot to update Dockerfile.
- Update `example_kind_test.go` to check for more details. Maybe you should run a client to verify what's needed.
- Use `task dev-build`, `task dev-build`, to deploy in your local kind cluster and test your changes.

## Usage

```bash
$ task
task: Available tasks for this project:
* clean:                          clean helm chart
* default:                        print this content
* dependencies-install-mac:       Install dependencies
* dev-build:                      Build all docker images and push to kind registry (change to kind context)         (aliases: build)
* dev-create:                     create a kind cluster (change to kind context)                                     (aliases: create)
* dev-delete:                     delete a kind cluster                                                              (aliases: delete)
* e2e-tests:                      Run the end-to-end and unit tests using go (change to kind temporary context)      (aliases: test)
* init:                           init go and helm chart (run only once)
* tidy:                           update dependencies and tidy go modules
```