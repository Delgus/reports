PKG := "github.com/delgus/reports"

.PHONY: all fmt lint dep testdb-up testdb-down test build clean help

all: help

fmt: ## gofmt all project
	@gofmt -l -s -w .

lint: ## Lint the files
	@golangci-lint run

dep: ## Get dependencies
	@go mod vendor

build: ## Build the binary file
	@go build -a -o bin/reporter -v $(PKG)/cmd/reporter

test: ## Run tests
	@go test -short -coverprofile=coverage.txt ./...

testdb-up: ## Up container with test database
	docker-compose -f testdb.yml up --build -d

testdb-down: ## Down container with test database
	docker-compose -f testdb.yml down --volume

clean: ## Remove previous build
	@rm -f bin/report1 bin/report2

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'