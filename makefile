PKG := "github.com/delgus/reports"

.PHONY: all fmt lint build clean help dep

all: fmt lint build test

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

dep: ## Get dependencies
	@go mod vendor

build: ## Build the binary file
	@go build -a -o bin/report1 -v $(PKG)/cmd/report1
	@go build -a -o bin/report2 -v $(PKG)/cmd/report2

test: ## Run integration tests
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

clean: ## Remove previous build
	@rm -f bin/report1 bin/report2

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'