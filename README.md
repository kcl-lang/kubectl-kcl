# Kubectl KCL Plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/kcl-lang/kubectl-kcl)](https://goreportcard.com/report/github.com/kcl-lang/kubectl-kcl)
[![GoDoc](https://godoc.org/github.com/kcl-lang/kubectl-kcl?status.svg)](https://godoc.org/github.com/kcl-lang/kubectl-kcl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kcl-lang/kubectl-kcl/blob/main/LICENSE)

[KCL](https://github.com/kcl-lang/kcl) is a constraint-based record & functional domain language. Full documents of KCL can be found [here](https://kcl-lang.io/).

This project is a `kubectl` plugin to generate, mutate and validate Kubernetes manifests using the KCL programming language.

## Installation

Use this as a `kubectl` plugin.

### From Krew Index

Add to `krew` index and install with:

```shell
kubectl krew index add kubectl-kcl https://github.com/kcl-lang/kubectl-kcl
kubectl krew install kubectl-kcl/kcl
```

### From GitHub Releases

Download the binary from GitHub releases, then copy the `kubectl-kcl` binary to your `PATH`. If not, you can also use the binary standalone.

## Usage

```shell
kubectl kcl run -f ./examples/kcl-run.yaml
```

## Developing

### Prerequisites

+ GoLang 1.21+

```shell
git clone https://github.com/kcl-lang/kubectl-kcl.git
cd kubectl-kcl
go run main.go
```

### Test

#### Unit Test

```shell
go test ./...
```

#### Integration Test

```shell
go run main.go run -f ./examples/kcl-run.yaml
```

## Guides for Developing KCL

Here's what you can do in the KCL script:

+ Read resources from `option("resource_list")`. The `option("resource_list")` complies with the [KRM Functions Specification](https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md#krm-functions-specification). You can read the input resources from `option("items")` and the `functionConfig` from `option("functionConfig")`.
+ Return a KRM list for output resources.
+ Return an error using `assert {condition}, {error_message}`.
+ Read the PATH variables. e.g. `option("PATH")`.
+ Read the environment variables. e.g. `option("env")`.

Full documents of KCL can be found [here](https://kcl-lang.io/).
