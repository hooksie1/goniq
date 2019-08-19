PROJECT_NAME := "goniq"
PKG := "gitlab.com/hooksie1/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage lint

all: build

lint: ## Lint the files
	@golint -set_exit_status $(PKG)

test: ## Run unittests
	@go test $(PKG)

coverage:
	@go test -cover $(PKG)

dep: ## Get the dependencies
	@go get -u golang.org/x/lint/golint
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'