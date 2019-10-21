PKG := "github.com/delgus/reports"

.PHONY: all fmt lint build clean help

all: build

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

build: ## Build the binary file
	@go build -a -o bin/report1 -v $(PKG)/cmd/report1
	@go build -a -o bin/report2 -v $(PKG)/cmd/report2

clean: ## Remove previous build
	@rm -f bin/report1 bin/report2

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'