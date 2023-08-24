.PHONY: default run build docs clean

# Variables
APP_NAME=opportunities-api-go

# Tasks
default: run

run:
	@go run ../cmd/opportunities/main.go
build:
	@go build -o $(APP_NAME) ../cmd/opportunities/main.go
docs:
	@swag init -g ../cmd/opportunities/main.go -o ../api
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./api