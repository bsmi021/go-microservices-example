.PHONY: help build test clean run-api run-mvc run-webserver docker-build docker-up docker-down lint

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build all services
	@echo "Building API service..."
	@go build -o bin/api ./src/api
	@echo "Building MVC service..."
	@go build -o bin/mvc ./mvc
	@echo "Building webserver..."
	@go build -o bin/webserver ./introduction/webserver
	@echo "Build complete!"

test: ## Run all tests
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

test-coverage: test ## Run tests and show coverage
	@go tool cover -html=coverage.txt

clean: ## Clean build artifacts
	@rm -rf bin/
	@rm -f coverage.txt
	@echo "Clean complete!"

run-api: ## Run API service
	@go run ./src/api

run-mvc: ## Run MVC service
	@go run ./mvc

run-webserver: ## Run simple webserver
	@go run ./introduction/webserver

docker-build: ## Build Docker images
	@docker-compose build

docker-up: ## Start all services with Docker Compose
	@docker-compose up

docker-down: ## Stop all services
	@docker-compose down

lint: ## Run linters
	@golangci-lint run

fmt: ## Format code
	@go fmt ./...

vet: ## Run go vet
	@go vet ./...

deps: ## Download dependencies
	@go mod download
	@go mod verify

tidy: ## Tidy go.mod
	@go mod tidy

.DEFAULT_GOAL := help
