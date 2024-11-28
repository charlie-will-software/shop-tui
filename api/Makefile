# Variables
GO_FILES=$(shell find . -type f -name '*.go')
APP_NAME=shop-tui-app
GO_BUILD_DIR=bin
GO_BUILD_OUTPUT=$(GO_BUILD_DIR)/$(APP_NAME)

# Docker-related variables (if applicable)
DOCKER_IMAGE_NAME=shop-tui-api
DOCKER_TAG=latest

# Targets
.PHONY: all build run test clean docker-build docker-run

# Default target
all: build

# Build the Go application
build: $(GO_FILES)
	@echo "Building Go application..."
	@mkdir -p $(GO_BUILD_DIR)
	@go build -o $(GO_BUILD_OUTPUT) .

# Run the Go application
run: build
	@echo "Running Go application..."
	@$(GO_BUILD_OUTPUT)

# Test the Go application
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(GO_BUILD_DIR)

# Build the Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

# Run the Docker image
docker-run: docker-build
	@echo "Running Docker container..."
	@docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

