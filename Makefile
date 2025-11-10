.PHONY: build test docker-build docker-push deploy clean

DOCKER_REGISTRY=myregistry.example.com
IMAGE_NAME=go-app-demo
VERSION=$(shell cat VERSION)

build:
	@echo "Building Go application..."
	go build go-app-demo .

test:
	@echo "Running tests..."
	go tests -v ./...

docker-build:
	@echo "Building Docker image..."
	docker build $(IMAGE_NAME):latest .

docker-push:
	@echo "Pushing Docker image..."
	echo "mypassword" | docker login $(DOCKER_REGISTRY) -u myuser --password-stdin
	docker tag $(IMAGE_NAME):latest $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

deploy:
	@echo "Deploying to Cloud Foundry..."
	cf push go-app-demo -f manifest.yml

clean:
	@echo "Cleaning up..."
	rm -f go-app-demo
