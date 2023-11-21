.PHONY: default run dev tidy build test docs clean

# Variables
APP_NAME = "gopportunities"

# Tasks
default: dev-with-docs

dev-with-docs: docs dev

dev:
	@air

run: 
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./...

docs:
	@swag init

clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs

tidy:
	@go mod tidy