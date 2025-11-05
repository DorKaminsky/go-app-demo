.PHONY: build test docker-build docker-push deploy clean

# ISSUE 1: Hardcoded values that should be variables
DOCKER_REGISTRY=myregistry.example.com
IMAGE_NAME=go-app-demo
VERSION=$(shell cat VERSION)

# ISSUE 2: No validation of required tools (docker, cf cli, go)

build:
	@echo "Building Go application..."
	go build -o go-app-demo .

test:
	@echo "Running tests..."
	go test -v ./...
	# ISSUE 3: No coverage report or coverage threshold check

# ISSUE 4: Missing lint target that CI pipeline expects
# lint:
# 	golangci-lint run

docker-build:
	@echo "Building Docker image..."
	# ISSUE 5: Using 'latest' tag only, no version tagging
	docker build -t $(IMAGE_NAME):latest .

docker-push:
	@echo "Pushing Docker image..."
	# ISSUE 6: Hardcoded registry credentials (security issue!)
	echo "mypassword" | docker login $(DOCKER_REGISTRY) -u myuser --password-stdin
	# ISSUE 7: Only pushing latest, not version tag
	docker tag $(IMAGE_NAME):latest $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

deploy:
	@echo "Deploying to Cloud Foundry..."
	# ISSUE 8: No check if cf CLI is installed or logged in
	# ISSUE 9: VERSION environment variable not normalized (still has -SNAPSHOT)
	cf push go-app-demo -f manifest.yml
	# ISSUE 10: No rollback mechanism or deployment verification

clean:
	@echo "Cleaning up..."
	rm -f go-app-demo
	# ISSUE 11: Doesn't clean Docker images or test artifacts
