BINARY_NAME=web-api

.PHONY: build run clean help

build:
	@echo "Building the application..."
	@go build -o bin/$(BINARY_NAME) cmd/web/main.go

run:
	@echo "Starting the application..."
	@./bin/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf bin

help:
	@echo "Available commands:"
	@echo "  build    Build the application"
	@echo "  run      Run the application"
	@echo "  clean    Remove the built binary"
