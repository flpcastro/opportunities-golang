.PHONY: default run build docs clean

# Variables
APP_NAME=opportunities-api-go

# Tasks
default: run-with-docs

run:
	@go run ../cmd/opportunities/main.go
run-with-docs:
	@swag init -d ../cmd/opportunities/ -o ../api
	@go run ../cmd/opportunities/main.go
build:
	@go build -o $(APP_NAME) ../cmd/opportunities/main.go
docs:
	@swag init -d ../cmd/opportunities/ -o ../api
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./api