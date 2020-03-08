PKG := "github.com/delgus/reports"

.PHONY: all fmt lint dep build test clean help

all: help

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

dep: ## Get dependencies
	@go mod vendor

build: ## Build the binary file
	@go build -a -o reporter -v $(PKG)/cmd/reporter

test: ## Run tests
	@go-acc -o coverage.txt ./...

clean: ## Remove previous build
	@rm -f reporter

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'