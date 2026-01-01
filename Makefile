.PHONY: build test test-coverage clean run install fmt vet lint help dev all check ci test-race

# Binary name
BINARY_NAME=spotify-toolbox
BUILD_DIR=bin

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/$(BUILD_DIR)

# Default target
.DEFAULT_GOAL := help

## help: Display this help message
help:
	@echo "Available targets:"
	@grep -E '^##' $(MAKEFILE_LIST) | sed 's/^## /  /'

## build: Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(GOBIN)
	@go build -o $(GOBIN)/$(BINARY_NAME) ./cmd/spotify-toolbox
	@echo "Build complete: $(GOBIN)/$(BINARY_NAME)"

## test: Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@go clean
	@echo "Clean complete"

## run: Build and run the application
run: build
	@$(GOBIN)/$(BINARY_NAME)

## install: Install dependencies
install:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed"

## fmt: Format Go code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

## lint: Run golangci-lint (requires golangci-lint to be installed)
lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/"; \
	fi

## dev: Install dependencies and build
dev: install build

## test-race: Run tests with race detector
test-race:
	@echo "Running tests with race detector..."
	@go test -v -race ./...

## check: Run all checks (fmt, vet, lint, test)
check: fmt vet lint test
	@echo "All checks passed!"

## ci: Run CI checks (used in GitHub Actions)
ci: check build
	@echo "CI checks complete!"

## all: Run fmt, vet, test, and build
all: fmt vet test build
