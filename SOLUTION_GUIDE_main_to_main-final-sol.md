# Solution Guide: main → main-final-sol

This document details ALL changes required to transform the **main** branch (with many basic issues) into the **main-final-sol** branch (production-ready code).

---

## 📊 Summary

The **main** branch has ~30+ issues ranging from basic syntax errors to security vulnerabilities and missing features. This guide shows how to fix them all.

| Category | Issues Fixed | Files Changed |
|----------|--------------|---------------|
| **Code Quality** | 6 | main.go, main_test.go |
| **Security** | 3 | Makefile, CI/CD |
| **Build Process** | 8 | Makefile, Dockerfile |
| **CI/CD** | 10 | .github/workflows/ci.yml |
| **Deployment** | 3 | manifest.yml |

**Total Changes:** ~30 fixes across 5 files

---

## 🔧 Complete Fix List

### **Category 1: Code Quality Fixes (main.go)**

#### **Fix #1: Replace Deprecated ioutil Package**
**File:** `main.go`
**Line:** 6

```diff
  import (
  	"encoding/json"
  	"fmt"
- 	"io/ioutil"
+ 	"context"
+ 	"log"
  	"net/http"
  	"os"
+ 	"os/signal"
  	"strings"
+ 	"syscall"
  	"time"
  )
```

**Reason:** `ioutil` is deprecated since Go 1.16. Use `os.ReadFile` instead.

---

#### **Fix #2: Add Version Normalization**
**File:** `main.go`
**Lines:** 18-30

```diff
  func getVersion() string {
  	version := os.Getenv("VERSION")
- 	if version != "" {
- 		return version
- 	}
-
- 	data, err := ioutil.ReadFile("VERSION")
+ 	if version == "" {
+ 		data, err := os.ReadFile("VERSION")
  		if err != nil {
  			return "unknown"
  		}
- 		return strings.TrimSpace(string(data))
+ 		version = strings.TrimSpace(string(data))
  	}
+
+ 	// Normalize version by stripping -SNAPSHOT suffix
+ 	version = strings.TrimSuffix(version, "-SNAPSHOT")
+ 	return version
  }
```

**Reason:** Version should strip `-SNAPSHOT` suffix for production deployments.

---

#### **Fix #3: Add HealthResponse Struct**
**File:** `main.go`
**After line 16**

```diff
  type InfoResponse struct {
  	Version    string `json:"version"`
  	DeployedAt string `json:"deployed_at"`
  }
+
+ type HealthResponse struct {
+ 	Status string `json:"status"`
+ }
```

**Reason:** Required for health check endpoint.

---

#### **Fix #4: Add Error Handling to infoHandler**
**File:** `main.go`
**Lines:** 32-40

```diff
  func infoHandler(w http.ResponseWriter, r *http.Request) {
  	response := InfoResponse{
  		Version:    getVersion(),
  		DeployedAt: time.Now().UTC().Format(time.RFC3339),
  	}

  	w.Header().Set("Content-Type", "application/json")
- 	json.NewEncoder(w).Encode(response)
+ 	if err := json.NewEncoder(w).Encode(response); err != nil {
+ 		log.Printf("Error encoding response: %v", err)
+ 		http.Error(w, "Internal server error", http.StatusInternalServerError)
+ 	}
  }
```

**Reason:** Proper error handling prevents silent failures.

---

#### **Fix #5: Add Health Check Endpoint**
**File:** `main.go`
**After infoHandler**

```diff
+ func healthHandler(w http.ResponseWriter, r *http.Request) {
+ 	response := HealthResponse{
+ 		Status: "ok",
+ 	}
+
+ 	w.Header().Set("Content-Type", "application/json")
+ 	if err := json.NewEncoder(w).Encode(response); err != nil {
+ 		log.Printf("Error encoding health response: %v", err)
+ 		http.Error(w, "Internal server error", http.StatusInternalServerError)
+ 	}
+ }
```

**Reason:** Required for Cloud Foundry and container orchestration health checks.

---

#### **Fix #6: Add Graceful Shutdown**
**File:** `main.go`
**Lines:** 42-52

```diff
  func main() {
  	port := os.Getenv("PORT")
  	if port == "" {
  		port = "8080"
  	}

+ 	// Register handlers
  	http.HandleFunc("/info", infoHandler)
+ 	http.HandleFunc("/health", healthHandler)

+ 	// Create server
+ 	server := &http.Server{
+ 		Addr:         ":" + port,
+ 		ReadTimeout:  15 * time.Second,
+ 		WriteTimeout: 15 * time.Second,
+ 		IdleTimeout:  60 * time.Second,
+ 	}
+
+ 	// Start server in goroutine
+ 	go func() {
+ 		fmt.Printf("Server starting on port %s\n", port)
+ 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
+ 			log.Fatalf("Server failed to start: %v", err)
+ 		}
+ 	}()
+
+ 	// Graceful shutdown
+ 	quit := make(chan os.Signal, 1)
+ 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
+ 	<-quit
+
+ 	log.Println("Server is shutting down...")
+
+ 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
+ 	defer cancel()
+
+ 	if err := server.Shutdown(ctx); err != nil {
+ 		log.Fatalf("Server forced to shutdown: %v", err)
+ 	}
+
+ 	log.Println("Server exited")
- 	fmt.Printf("Server starting on port %s\n", port)
- 	http.ListenAndServe(":"+port, nil)
  }
```

**Reason:** Production applications must handle shutdown signals gracefully to avoid dropping in-flight requests.

---

### **Category 2: Testing Fixes (main_test.go)**

#### **Fix #7: Complete Test Rewrite**
**File:** `main_test.go`

The original tests are minimal. Replace entire file with comprehensive tests:

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Test version normalization
	os.Setenv("VERSION", "1.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	version := getVersion()

	// Should strip -SNAPSHOT
	if version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", version)
	}
}

func TestGetVersionFromFile(t *testing.T) {
	// Clear env var to test file reading
	os.Unsetenv("VERSION")

	// Create temporary VERSION file
	content := []byte("2.0.0-SNAPSHOT\n")
	err := os.WriteFile("VERSION", content, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("VERSION")

	version := getVersion()

	// Should read from file and strip -SNAPSHOT
	if version != "2.0.0" {
		t.Errorf("Expected version '2.0.0', got '%s'", version)
	}
}

func TestInfoHandler(t *testing.T) {
	os.Setenv("VERSION", "1.0.0")
	defer os.Unsetenv("VERSION")

	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response InfoResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", response.Version)
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response HealthResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}
```

**Changes:**
- Added test for version normalization
- Added test for file reading
- Added health endpoint test
- Added proper assertions
- Added content-type checks

---

### **Category 3: Build Process Fixes (Makefile)**

#### **Fix #8: Complete Makefile Rewrite**
**File:** `Makefile`

Replace entire file:

```makefile
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
```

**Key Changes:**
- ❌ Removed hardcoded `DOCKER_REGISTRY=myregistry.example.com`
- ❌ Removed hardcoded password in docker-push
- ✅ Added `-o` flag to build command
- ✅ Fixed `go tests` → `go test`
- ✅ Fixed `docker build` syntax (added `-t`)
- ✅ Added version normalization logic
- ✅ Added `check-tools`, `lint`, `coverage` targets
- ✅ Uses environment variables for Docker credentials
- ✅ Proper error checking

---

### **Category 4: Docker Fixes (Dockerfile)**

#### **Fix #9: Complete Dockerfile Rewrite**
**File:** `Dockerfile`

Replace entire file:

```dockerfile
# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod ./
RUN go mod download

# Copy source code
COPY main.go ./
COPY VERSION ./

# Build application with optimization flags
RUN go build -ldflags="-w -s" -o go-app-demo .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Create non-root user for security
RUN adduser -D -u 1000 appuser

# Copy binary and VERSION file from builder
COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port
ENV PORT=8080
EXPOSE ${PORT}

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

# Run application
CMD ["./go-app-demo"]
```

**Key Changes:**
- ✅ Fixed Go version from 1.21 → 1.22
- ✅ Added WORKDIR directives
- ✅ Optimized layer caching (copy go.mod first)
- ✅ Added non-root user (appuser)
- ✅ Added HEALTHCHECK directive
- ✅ Added optimization flags (-ldflags)
- ✅ Multi-stage build using alpine

---

### **Category 5: CI/CD Fixes (.github/workflows/ci.yml)**

#### **Fix #10: Complete CI/CD Rewrite**
**File:** `.github/workflows/ci.yml`

Replace entire file:

```yaml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

env:
  GO_VERSION: '1.22'
  IMAGE_NAME: go-app-demo

jobs:
  build:
    name: Build Application
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Build
        run: make build

  test:
    name: Run Tests
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Run tests
        run: make test

      - name: Run coverage
        run: make coverage

  lint:
    name: Lint Code
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ env.GO_VERSION }}
          cache: true

      - name: Run linter
        run: make lint

  docker-build:
    name: Build Docker Image
    runs-on: ubuntu-latest
    needs: [test, lint]
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Build Docker image
        uses: docker/build-push-action@v5
        with:
          context: .
          push: false
          tags: ${{ env.IMAGE_NAME }}:latest
          cache-from: type=gha
          cache-to: type=gha,mode=max

  docker-push:
    name: Push Docker Image
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    needs: [docker-build]
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ vars.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Extract version
        id: version
        run: |
          RAW_VERSION=$(cat VERSION)
          VERSION=$(echo $RAW_VERSION | sed 's/-SNAPSHOT//')
          echo "version=$VERSION" >> $GITHUB_OUTPUT

      - name: Build and push Docker image
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            ${{ vars.DOCKERHUB_USERNAME }}/${{ env.IMAGE_NAME }}:latest
            ${{ vars.DOCKERHUB_USERNAME }}/${{ env.IMAGE_NAME }}:${{ steps.version.outputs.version }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

**Key Changes:**
- ❌ Removed deploy job (CF deployment)
- ✅ Changed runners from `[self-hosted, solinas]` → `ubuntu-latest`
- ✅ Fixed Go version from 1.21 → 1.22
- ✅ Added env variables for GO_VERSION and IMAGE_NAME
- ✅ Added proper job dependencies (`needs:`)
- ✅ Added Docker caching (cache-from, cache-to)
- ✅ Added coverage step
- ✅ Uses GitHub Secrets for Docker credentials
- ✅ Removed hardcoded CF credentials

---

### **Category 6: Deployment Fixes (manifest.yml)**

#### **Fix #11: Complete manifest.yml Rewrite**
**File:** `manifest.yml`

Replace entire file:

```yaml
---
applications:
  - name: go-app-demo
    memory: 256M
    disk_quota: 512M
    instances: 2

    # Use Docker image from Docker Hub
    docker:
      image: ${DOCKERHUB_USERNAME}/go-app-demo:latest

    # VERSION is normalized by CI/CD pipeline (strips -SNAPSHOT)
    env:
      VERSION: 1.189.0

    # Health check configuration
    health-check-type: http
    health-check-http-endpoint: /health
    health-check-invocation-timeout: 5

    # Route configuration - let CF auto-generate the route
    random-route: true

    # Resource limits and scaling
    processes:
      - type: web
        instances: 2
        memory: 256M
        disk_quota: 512M

    # Timeout settings
    timeout: 60
```

**Key Changes:**
- ✅ Changed from buildpacks → docker deployment
- ✅ Added health check configuration
- ✅ Added resource limits
- ✅ Added proper VERSION normalization
- ✅ Added route configuration
- ✅ Uses Docker Hub image

---

## ✅ Summary of All Changes

### **Security Fixes:**
1. ❌ Removed hardcoded password from Makefile
2. ❌ Removed hardcoded Docker registry
3. ❌ Removed hardcoded CF credentials from CI/CD
4. ✅ Added non-root Docker user

### **Code Quality:**
5. ✅ Replaced deprecated ioutil
6. ✅ Added version normalization
7. ✅ Added health endpoint
8. ✅ Added graceful shutdown
9. ✅ Added proper error handling

### **Build & Testing:**
10. ✅ Fixed missing `-o` flag
11. ✅ Fixed `go tests` typo
12. ✅ Fixed `docker build` syntax
13. ✅ Added lint target
14. ✅ Added coverage target
15. ✅ Comprehensive test suite

### **Infrastructure:**
16. ✅ Fixed Go version 1.21 → 1.22
17. ✅ Added WORKDIR to Dockerfile
18. ✅ Optimized Docker layer caching
19. ✅ Added Docker HEALTHCHECK
20. ✅ Changed runners to ubuntu-latest
21. ✅ Added job dependencies
22. ✅ Added GitHub Actions caching

### **Deployment:**
23. ✅ Changed manifest from buildpack → docker
24. ✅ Added health check configuration
25. ✅ Added proper resource limits

---

## 🎯 Verification

After applying all fixes:

```bash
# Build succeeds
make build

# Tests pass
make test

# Coverage report
make coverage

# Linter passes
make lint

# Docker builds
make docker-build

# Run locally
docker run -p 8080:8080 go-app-demo:latest

# Test endpoints
curl http://localhost:8080/info
curl http://localhost:8080/health
```

---

*Generated with Claude Code*
