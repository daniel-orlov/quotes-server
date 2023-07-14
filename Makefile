.DEFAULT_GOAL := help

# This points to the repository root directory
repo_root = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: run-server
run-server: ## Run the server passing the arguments
	@echo "> Running the server ..."
	go run cmd/server/main.go $(ARGS)

.PHONY: run-client
run-client: ## Run the client passing the arguments
	@echo "> Running the client ..."
	go run cmd/client/main.go $(ARGS)

.PHONY: test
test: ## Run tests with coverage
	@echo "> Testing..."
	go clean --testcache
	go test -v ./... -coverprofile=$(repo_root)/test_coverage.out &&\
	echo "total coverage: $$(go tool cover -func=$(repo_root)/test_coverage.out | grep total: | awk '{ print $$3}')"\

.PHONY: tidy
tidy: ## Clean and format Go code
	@echo "> Tidying..."
	go mod tidy
	go fmt ./...
	@echo "> Done!"

.PHONY: fmt
fmt: ## Format Go code
	go fmt ./...

.PHONY: lint-host
lint-host: ## Run golangci-lint directly on host
	@echo "> Linting..."
	golangci-lint run -c .golangci.yml -v
	@echo "> Done!"

.PHONY: start
start: ## Run the server and client
	@echo "> Running the server and client ..."
	cd deploy && \
	docker-compose up \
	--abort-on-container-exit \
	--force-recreate \
	--pull always

.PHONY: help
help: ## Show this help
	@echo "make run-server - Run the server passing the arguments"
	@echo "make run-client - Run the client passing the arguments"
	@echo "make start - Run the server and client using docker-compose"
	@echo "make test - Run tests"
	@echo "make tidy - Clean and format Go code"
	@echo "make fmt - Format Go code"
	@echo "make lint-host - Run golangci-lint directly on host"
	@echo "make help - Show this help"
