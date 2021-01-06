-include .makerc

.PHONY: help build-local run-local build ci release publish

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build application
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/service/app cmd/service/main.go

run: ## Run application on local
	@go run cmd/service/main.go

clean:
	@rm -rf bin

swagger:
	SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ./api/v1/local/feed.yaml
