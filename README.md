# Kubectl KCL Plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/kcl-lang/kubectl-kcl)](https://goreportcard.com/report/github.com/kcl-lang/kubectl-kcl)
[![GoDoc](https://godoc.org/github.com/kcl-lang/kubectl-kcl?status.svg)](https://godoc.org/github.com/kcl-lang/kubectl-kcl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kcl-lang/kubectl-kcl/blob/main/LICENSE)

[KCL](https://github.com/kcl-lang/kcl) is a constraint-based record & functional domain language. Full documents of KCL can be found [here](https://kcl-lang.io/).

This project is a `kubectl` plugin to generate, mutate and validate Kubernetes manifests using the KCL programming language.

## Installation

### Krew

Add to `krew` index and install with:

```shell
kubectl krew index add kubectl-kcl https://github.com/kcl-lang/kubectl-kcl
kubectl krew install kubectl-kcl/kubectl-kcl
```

### GitHub release

Download the binary from GitHub releases.

If you want to use this as a `kubectl` plugin, then copy the `kubectl-kcl` binary to your `PATH`. If not, you can also use the binary standalone.

## Build

### Prerequisites

+ GoLang 1.19+

```shell
git clone https://github.com/kcl-lang/kubectl-kcl.git
cd kubectl-kcl
go run main.go
```

## Test

### Unit Test

```shell
go test -v ./...
```

### Integration Test

```bash
go run main.go run --file ./examples/kcl-run.yaml
```

## Guides for Developing KCL

Here's what you can do in the KCL script:

+ Read resources from `option("resource_list")`. The `option("resource_list")` complies with the [KRM Functions Specification](https://kpt.dev/book/05-developing-functions/01-functions-specification). You can read the input resources from `option("resource_list")["items"]` and the `functionConfig` from `option("resource_list")["functionConfig"]`.
+ Return a KPM list for output resources.
+ Return an error using `assert {condition}, {error_message}`.
+ Read the environment variables. e.g. `option("PATH")` (Not yet implemented).
+ Read the OpenAPI schema. e.g. `option("open_api")["definitions"]["io.k8s.api.apps.v1.Deployment"]` (Not yet implemented).

Full documents of KCL can be found [here](https://kcl-lang.io/).
