.PHONY: build test coverage lint docker-build docker-push deploy clean

DOCKER_REGISTRY ?= $(shell echo $$DOCKER_REGISTRY)
DOCKER_USERNAME ?= $(shell echo $$DOCKER_USERNAME)
DOCKER_PASSWORD ?= $(shell echo $$DOCKER_PASSWORD)
IMAGE_NAME=go-app-demo
VERSION=$(shell cat VERSION)

build:
	@echo "Building Go application..."
	go build -o go-app-demo .

test:
	@echo "Running tests..."
	go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed, using go vet..."; \
		go vet ./...; \
		go fmt ./...; \
	fi

docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):latest .
	docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(VERSION)

docker-push:
	@echo "Pushing Docker image..."
	@if [ -z "$(DOCKER_REGISTRY)" ] || [ -z "$(DOCKER_USERNAME)" ] || [ -z "$(DOCKER_PASSWORD)" ]; then \
		echo "Error: DOCKER_REGISTRY, DOCKER_USERNAME, and DOCKER_PASSWORD environment variables must be set"; \
		exit 1; \
	fi
	@echo "$$DOCKER_PASSWORD" | docker login $(DOCKER_REGISTRY) -u $(DOCKER_USERNAME) --password-stdin
	docker tag $(IMAGE_NAME):latest $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(VERSION)
	docker tag $(IMAGE_NAME):latest $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

deploy:
	@echo "Deploying to Cloud Foundry..."
	@if [ -z "$(CF_API)" ] || [ -z "$(CF_USERNAME)" ] || [ -z "$(CF_PASSWORD)" ] || [ -z "$(CF_ORG)" ] || [ -z "$(CF_SPACE)" ]; then \
		echo "Error: CF_API, CF_USERNAME, CF_PASSWORD, CF_ORG, and CF_SPACE environment variables must be set"; \
		exit 1; \
	fi
	cf login -a $(CF_API) -u $(CF_USERNAME) -p $(CF_PASSWORD) -o $(CF_ORG) -s $(CF_SPACE)
	cf push go-app-demo -f manifest.yml

clean:
	@echo "Cleaning up..."
	rm -f go-app-demo
	rm -f coverage.out coverage.html