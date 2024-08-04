# Variables
IMAGE_NAME = shop-tui
VERSION = latest

# Build the Docker container
build:
	docker build -t $(IMAGE_NAME):$(VERSION) .

.PHONY: build
