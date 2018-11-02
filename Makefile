GO_VERSION ?= 1.11
GOOS ?= darwin
GOARCH ?= amd64
TAG := $(shell git rev-parse --short HEAD)-go${GO_VERSION}

FILEEXT :=
ifeq (${GOOS},windows)
FILEEXT := .exe
endif
CUSTOMTAG ?=

.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo help

##@ Dependencies

#.PHONY: build-dev-deps:
#build-dev-deps: ## Install dependencies for builds
#	go get github.com/mattn/goveralls
#	go get golang.org/x/tools/cover
#	go get github.com/modocache/gover
#	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin v1.11

.PHONY: lint
lint: ## Lint
	@echo "Running ${@}"
	@golangci-lint run

.PHONY: test
test: ## Run tests
	go test -race -v -cover -coverprofile=.coverprofile ./...
	@echo done

.PHONY: build
build: ## Build binaries
	mkdir -p builds
	go build -o builds/kubesync-${TAG}${FILEEXT} .
