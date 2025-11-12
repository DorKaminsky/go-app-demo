.PHONY: build test lint coverage docker-build docker-push deploy clean check-tools

# Configuration - can be overridden via environment variables
IMAGE_NAME ?= go-app-demo
RAW_VERSION := $(shell cat VERSION 2>/dev/null || echo "unknown")
VERSION := $(shell echo $(RAW_VERSION) | sed 's/-SNAPSHOT//')
GIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Check if required tools are installed
check-tools:
	@command -v go >/dev/null 2>&1 || { echo "Error: go is not installed"; exit 1; }
	@command -v docker >/dev/null 2>&1 || { echo "Error: docker is not installed"; exit 1; }
	@echo "✓ Required tools are installed"

build: check-tools
	@echo "Building Go application..."
	@go build -o go-app-demo .
	@echo "✓ Build complete"

test: check-tools
	@echo "Running tests..."
	@go test -v ./...
	@echo "✓ Tests passed"

coverage: check-tools
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@echo "✓ Coverage report generated"

lint: check-tools
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "Warning: golangci-lint not installed, running go vet instead"; \
		go vet ./...; \
	fi
	@echo "✓ Lint complete"

docker-build: check-tools
	@echo "Building Docker image..."
	@docker build \
		-t $(IMAGE_NAME):latest \
		-t $(IMAGE_NAME):$(VERSION) \
		-t $(IMAGE_NAME):$(VERSION)-$(GIT_SHA) \
		.
	@echo "✓ Docker image built with tags: latest, $(VERSION), $(VERSION)-$(GIT_SHA)"

docker-push: docker-build
	@echo "Pushing Docker image..."
	@if [ -z "$$DOCKERHUB_USERNAME" ] || [ -z "$$DOCKERHUB_TOKEN" ]; then \
		echo "Error: DOCKERHUB_USERNAME and DOCKERHUB_TOKEN must be set"; \
		exit 1; \
	fi
	@echo "$$DOCKERHUB_TOKEN" | docker login -u "$$DOCKERHUB_USERNAME" --password-stdin
	@docker tag $(IMAGE_NAME):latest $$DOCKERHUB_USERNAME/$(IMAGE_NAME):latest
	@docker tag $(IMAGE_NAME):$(VERSION) $$DOCKERHUB_USERNAME/$(IMAGE_NAME):$(VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION)-$(GIT_SHA) $$DOCKERHUB_USERNAME/$(IMAGE_NAME):$(VERSION)-$(GIT_SHA)
	@docker push $$DOCKERHUB_USERNAME/$(IMAGE_NAME):latest
	@docker push $$DOCKERHUB_USERNAME/$(IMAGE_NAME):$(VERSION)
	@docker push $$DOCKERHUB_USERNAME/$(IMAGE_NAME):$(VERSION)-$(GIT_SHA)
	@echo "✓ Images pushed: latest, $(VERSION), $(VERSION)-$(GIT_SHA)"

deploy:
	@echo "Deploying to Cloud Foundry..."
	@command -v cf >/dev/null 2>&1 || { echo "Error: cf CLI is not installed"; exit 1; }
	@cf target >/dev/null 2>&1 || { echo "Error: Not logged in to Cloud Foundry"; exit 1; }
	@echo "Normalizing VERSION to $(VERSION) (stripped -SNAPSHOT)"
	@sed -i.bak 's/VERSION:.*/VERSION: $(VERSION)/' manifest.yml && rm manifest.yml.bak || sed -i '' 's/VERSION:.*/VERSION: $(VERSION)/' manifest.yml
	@cf push go-app-demo -f manifest.yml
	@echo "Verifying deployment..."
	@sleep 5
	@cf app go-app-demo
	@echo "✓ Deployment complete"

clean:
	@echo "Cleaning up..."
	@rm -f go-app-demo
	@rm -f coverage.out
	@rm -f *.test
	@docker rmi $(IMAGE_NAME):latest $(IMAGE_NAME):$(VERSION) 2>/dev/null || true
	@echo "✓ Cleanup complete"
