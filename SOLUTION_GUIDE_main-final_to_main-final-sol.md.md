# Solution Guide: main-final → main-final-sol

This document details all changes made to fix the 8 intentional DevOps issues.

---

## 📊 Summary

| Issue | File | Lines Changed | Category |
|-------|------|---------------|----------|
| #1 | Dockerfile | 2, 20 | Configuration |
| #2 | Makefile | 17 | Build Process |
| #3 | Makefile | 22 | Build Process |
| #4 | main_test.go | 39-41 | Testing |
| #5 | manifest.yml | 8-10 | Deployment |
| #6 | .github/workflows/ci.yml | 18, 34, 53, 69, 89 | CI/CD |
| #7 | .github/workflows/ci.yml | 70, 91 | CI/CD |
| #8 | .github/workflows/ci.yml | 21, 24, 37, 40, 56, 59, 73, 76, 79, 94, 97, 100, 113 | CI/CD |

**Total Files Changed:** 5
**Total Lines Modified:** ~40

---

## 🔧 Detailed Changes

### **Issue #1: Missing WORKDIR in Dockerfile**

**Problem:** Both build and runtime stages were missing `WORKDIR /app` directives, causing files to be copied to wrong locations.

**File:** `Dockerfile`

**Changes:**

**Line 4 - Build stage:**
```diff
  # Build stage
  FROM golang:1.22-alpine AS builder

+ WORKDIR /app
+
  # Copy go mod files first for better layer caching
  COPY go.mod ./
```

**Line 20 - Runtime stage:**
```diff
  # Runtime stage
  FROM alpine:latest

+ WORKDIR /app
+
  # Create non-root user for security
  RUN adduser -D -u 1000 appuser
```

**Impact:** Without WORKDIR, files would be copied to `/` instead of `/app`, breaking the application.

---

### **Issue #2: Missing -o Flag in go build**

**Problem:** The `go build` command was missing the `-o` flag to specify output binary name.

**File:** `Makefile`

**Change - Line 17:**
```diff
  build: check-tools
  	@echo "Building Go application..."
- 	@go build go-app-demo .
+ 	@go build -o go-app-demo .
  	@echo "✓ Build complete"
```

**Impact:** Without `-o`, the build would fail with "cannot use non-main package" error because Go interprets `go-app-demo` as a package name instead of output filename.

---

### **Issue #3: Typo in Makefile (go tests)**

**Problem:** Command `go tests` doesn't exist; should be `go test`.

**File:** `Makefile`

**Change - Line 22:**
```diff
  test: check-tools
  	@echo "Running tests..."
- 	@go tests -v ./...
+ 	@go test -v ./...
  	@echo "✓ Tests passed"
```

**Impact:** Running `make test` would fail with "unknown command 'tests'" error.

---

### **Issue #4: Failing Test - Wrong Expected Value**

**Problem:** Test expected version "3.0.0" but function returns "2.0.0", causing test failure.

**File:** `main_test.go`

**Change - Lines 39-41:**
```diff
  	version := getVersion()

- 	// This test will fail because it expects wrong value
- 	if version != "3.0.0" {
- 		t.Errorf("Expected version '3.0.0', got '%s'", version)
+ 	// Should read from file and strip -SNAPSHOT
+ 	if version != "2.0.0" {
+ 		t.Errorf("Expected version '2.0.0', got '%s'", version)
  	}
```

**Impact:** Test suite would fail, blocking CI/CD pipeline.

---

### **Issue #5: manifest.yml Using Buildpack Instead of Docker**

**Problem:** Cloud Foundry manifest configured to use buildpack, but the workflow builds and pushes Docker images.

**File:** `manifest.yml`

**Change - Lines 8-10:**
```diff
      memory: 256M
      disk_quota: 512M
      instances: 2

-     # Use Go buildpack
-     buildpacks:
-       - go_buildpack
+     # Use Docker image from Docker Hub
+     docker:
+       image: ${DOCKERHUB_USERNAME}/go-app-demo:latest

      # VERSION is normalized by CI/CD pipeline (strips -SNAPSHOT)
      env:
```

**Impact:** Deployment would fail because the pipeline pushes Docker images but manifest expects buildpack deployment.

---

### **Issue #6: Self-Hosted Runners Instead of ubuntu-latest**

**Problem:** All jobs configured to use `[self-hosted, solinas]` runners that candidates don't have access to.

**File:** `.github/workflows/ci.yml`

**Changes:**

**Line 18 - build job:**
```diff
    build:
      name: Build Application
-     runs-on: [ self-hosted, solinas ]
+     runs-on: ubuntu-latest
```

**Line 34 - test job:**
```diff
    test:
      name: Run Tests
-     runs-on: [ self-hosted, solinas ]
+     runs-on: ubuntu-latest
```

**Line 53 - lint job:**
```diff
    lint:
      name: Lint Code
-     runs-on: [ self-hosted, solinas ]
+     runs-on: ubuntu-latest
```

**Line 69 - docker-build job:**
```diff
    docker-build:
      name: Build Docker Image
-     runs-on: [ self-hosted, solinas ]
+     runs-on: ubuntu-latest
```

**Line 89 - docker-push job:**
```diff
    docker-push:
      name: Push Docker Image
-     runs-on: [ self-hosted, solinas ]
+     runs-on: ubuntu-latest
```

**Impact:** CI/CD pipeline would never start because the self-hosted runners don't exist in the candidate's fork.

---

### **Issue #7: Typo 'need' Instead of 'needs'**

**Problem:** GitHub Actions uses `needs:` keyword (plural), not `need:`.

**File:** `.github/workflows/ci.yml`

**Changes:**

**Line 70 - docker-build job:**
```diff
    docker-build:
      name: Build Docker Image
      runs-on: ubuntu-latest
-     need: [test, lint]
+     needs: [test, lint]
      steps:
```

**Line 91 - docker-push job:**
```diff
    docker-push:
      name: Push Docker Image
      runs-on: ubuntu-latest
      if: github.ref == 'refs/heads/main'
-     need: [docker-build]
+     needs: [docker-build]
      steps:
```

**Impact:** GitHub Actions would fail with "unexpected value 'need'" error, and jobs would run in parallel instead of sequentially.

---

### **Issue #8: Deprecated GitHub Actions Versions**

**Problem:** Using older versions of actions (@v2, @v3, @v4) when newer versions are available.

**File:** `.github/workflows/ci.yml`

**Changes:**

**Build job (lines 21, 24):**
```diff
        steps:
          - name: Checkout code
-           uses: actions/checkout@v2
+           uses: actions/checkout@v3

          - name: Set up Go
-           uses: actions/setup-go@v3
+           uses: actions/setup-go@v4
```

**Test job (lines 37, 40):**
```diff
        steps:
          - name: Checkout code
-           uses: actions/checkout@v2
+           uses: actions/checkout@v3

          - name: Set up Go
-           uses: actions/setup-go@v3
+           uses: actions/setup-go@v4
```

**Lint job (lines 56, 59):**
```diff
        steps:
          - name: Checkout code
-           uses: actions/checkout@v2
+           uses: actions/checkout@v3

          - name: Set up Go
-           uses: actions/setup-go@v3
+           uses: actions/setup-go@v4
```

**Docker-build job (lines 73, 76, 79):**
```diff
        steps:
          - name: Checkout code
-           uses: actions/checkout@v2
+           uses: actions/checkout@v3

          - name: Set up Docker Buildx
-           uses: docker/setup-buildx-action@v2
+           uses: docker/setup-buildx-action@v3

          - name: Build Docker image
-           uses: docker/build-push-action@v4
+           uses: docker/build-push-action@v5
```

**Docker-push job (lines 94, 97, 100, 113):**
```diff
        steps:
          - name: Checkout code
-           uses: actions/checkout@v2
+           uses: actions/checkout@v3

          - name: Set up Docker Buildx
-           uses: docker/setup-buildx-action@v2
+           uses: docker/setup-buildx-action@v3

          - name: Log in to Docker Hub
-           uses: docker/login-action@v2
+           uses: docker/login-action@v3

          ...

          - name: Build and push Docker image
-           uses: docker/build-push-action@v4
+           uses: docker/build-push-action@v5
```

**Impact:**
- Deprecated actions may have security vulnerabilities
- Missing newer features and performance improvements
- GitHub will show deprecation warnings
- May stop working when GitHub removes old versions

---

## ✅ Verification

After applying all fixes, the following should work:

```bash
# Build succeeds
make build

# Tests pass
make test

# Docker builds
make docker-build

# Docker image works
docker run -p 8080:8080 go-app-demo:latest

# CI/CD pipeline runs (in GitHub Actions)
git push origin main-final-sol
```

---

## 📈 Learning Outcomes

Candidates who successfully fix these issues demonstrate:

1. **Docker Knowledge:** Understanding WORKDIR and its importance
2. **Makefile/Build Skills:** Proper command syntax and flags
3. **Testing Best Practices:** Reading and fixing test expectations
4. **Cloud Foundry Understanding:** Docker vs buildpack deployments
5. **CI/CD Expertise:** GitHub Actions syntax, runners, and dependencies
6. **Attention to Detail:** Catching typos and version inconsistencies
7. **DevOps Maturity:** Understanding security implications of deprecated dependencies

---

## 🎯 Candidate Evaluation Criteria

### Must Fix (Critical - 5 issues)
- ✅ Issue #2: Missing -o flag (blocks build)
- ✅ Issue #3: go tests typo (blocks testing)
- ✅ Issue #6: Self-hosted runners (blocks CI/CD)
- ✅ Issue #7: needs typo (breaks job dependencies)
- ✅ Issue #5: manifest.yml mismatch (blocks deployment)

### Should Fix (Important - 2 issues)
- ✅ Issue #1: Missing WORKDIR (Docker fails)
- ✅ Issue #4: Failing test (quality gates)

### Nice to Fix (Best Practice - 1 issue)
- ✅ Issue #8: Deprecated actions (security/maintenance)

**Target:** Candidates should fix at least 6-7 out of 8 issues to pass.

---

*Generated with Claude Code*
