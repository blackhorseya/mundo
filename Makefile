VERSION := $(shell git describe --tags --always)

## common
.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## show version
	@echo $(VERSION)

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out bin
	@echo Successfuly removed artifacts

## go
.PHONY: lint
lint: ## run golangci-lint
	@golangci-lint run ./...

.PHONY: gazelle
gazelle: ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build
build: ## build go binary
	@bazel build //...

.PHONY: test
test: ## test go binary
	@bazel test //...

.PHONY: docker-push
docker-push: ## push docker image
	@bazel run //adapter:push -- --tag=$(VERSION)
