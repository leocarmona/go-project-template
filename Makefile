PWD = $(shell pwd -L)
GOCMD=go
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose
GOTEST=$(GOCMD) test
IMAGE_NAME = go-project-template
LAMBDA_BINARY = $(IMAGE_NAME)-lambda-binary
LIBRARY_ENV ?= dev

.PHONY: all test vendor

all: help

help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build image
	$(DOCKERCMD) build -t $(IMAGE_NAME) .

build-lambda: lint ## Build lambda binary
	rm -rfv $(LAMBDA_BINARY);
	GOOS=linux GOARCH=amd64 go build -o $(LAMBDA_BINARY) ./cmd/main.go;
	zip -9 $(LAMBDA_BINARY).zip $(LAMBDA_BINARY);
	rm -rfv $(LAMBDA_BINARY);

run: ## Run API
	$(DOCKERCMD) run --net=host --env DB_PORT=6432 -d -t $(IMAGE_NAME)

stop: ## Stop API
	$(DOCKERCMD) stop $$($(DOCKERCMD) ps -q --filter ancestor=go-project-template )

compose-up: ## Run docker-compose services of project
	$(DOCKERCOMPOSECMD) up -d

compose-down: ## Stop docker-compose services of project
	$(DOCKERCOMPOSECMD) down --remove-orphans

compose-restart: compose-down compose-up ## Restart docker-compose services of project

compose-logs: ## Logs docker-compose containers of project
	$(DOCKERCOMPOSECMD) logs -f app

lint: lint-go lint-yaml ## Run all available linters

lint-go: fmt ## Use golintci-lint on your project
	$(DOCKERCMD) run --rm -v $(PWD):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s

lint-yaml: fmt ## Use yamllint on the yaml file of your projects
	$(DOCKERCMD) run --rm $$(tty -s && echo "-it" || echo) -v $(PWD):/data cytopia/yamllint:latest .

fmt: tidy ## Run go fmt
	$(GOCMD) fmt ./...

tidy: ## Downloads go dependencies
	$(GOCMD) mod tidy
