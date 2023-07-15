VERSION := $(shell cat VERSION)

PKG:= kcl-lang.io/kubectl-kcl
LDFLAGS := -X $(PKG)/cmd.Version=$(VERSION)

GO ?= go

.PHONY: run
run:
	go run main.go run -f ./examples/kcl-run.yaml
	go run main.go apply -f ./examples/kcl-apply.yaml

.PHONY: format
format:
	test -z "$$(find . -type f -o -name '*.go' -exec gofmt -d {} + | tee /dev/stderr)" || \
	test -z "$$(find . -type f -o -name '*.go' -exec gofmt -w {} + | tee /dev/stderr)"

.PHONY: lint
lint:
	scripts/update-gofmt.sh
	scripts/verify-gofmt.sh
	# scripts/verify-golint.sh
	scripts/verify-govet.sh

.PHONY: build
build: lint
	mkdir -p bin/
	go build -v -o bin/kcl -ldflags="$(LDFLAGS)"

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: bootstrap
bootstrap:
	go mod download
	command -v golint || GO111MODULE=off go get -u golang.org/x/lint/golint

.PHONY: docker-run-release
docker-run-release: export pkg=/go/src/github.com/kcl-lang/kubectl-kcl
docker-run-release:
	git checkout main
	git push
	docker run -it --rm -e GITHUB_TOKEN -v $(shell pwd):$(pkg) -w $(pkg) golang:1.19.1 make bootstrap release

.PHONY: dist
dist: export COPYFILE_DISABLE=1 #teach OSX tar to not put ._* files in tar archive
dist: export CGO_ENABLED=0
dist:
	rm -rf build/kubectl-kcl/* release/*
	mkdir -p build/kubectl-kcl/bin release/
	cp -f README.md LICENSE build/kubectl-kcl
	GOOS=linux GOARCH=amd64 $(GO) build -o build/kubectl-kcl/bin/kubectl-kcl -trimpath -ldflags="$(LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/kubectl-kcl-linux-amd64.tgz kubectl-kcl/
	GOOS=linux GOARCH=arm64 $(GO) build -o build/kubectl-kcl/bin/kubectl-kcl -trimpath -ldflags="$(LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/kubectl-kcl-linux-arm64.tgz kubectl-kcl/
	GOOS=darwin GOARCH=amd64 $(GO) build -o build/kubectl-kcl/bin/kubectl-kcl -trimpath -ldflags="$(LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/kubectl-kcl-macos-amd64.tgz kubectl-kcl/
	GOOS=darwin GOARCH=arm64 $(GO) build -o build/kubectl-kcl/bin/kubectl-kcl -trimpath -ldflags="$(LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/kubectl-kcl-macos-arm64.tgz kubectl-kcl/
	rm build/kubectl-kcl/bin/kubectl-kcl
	GOOS=windows GOARCH=amd64 $(GO) build -o build/kubectl-kcl/bin/kubectl-kcl.exe -trimpath -ldflags="$(LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/kubectl-kcl-windows-amd64.tgz kubectl-kcl/

.PHONY: release
release: lint dist
	scripts/release.sh v$(VERSION)
