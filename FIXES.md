# DevOps Issues - Fixed

This document details all the fixes applied to resolve the intentional DevOps issues in the repository.

## Summary

- **Total Issues Fixed**: 52+
- **Files Modified**: 6 (main.go, main_test.go, Dockerfile, Makefile, .github/workflows/ci.yml, manifest.yml)
- **New Files Created**: 1 (FIXES.md)

---

## 1. Go Code Fixes (`main.go`)

### ✅ ISSUE 1: Using deprecated `ioutil` package
**Fix**: Replaced `ioutil.ReadFile()` with `os.ReadFile()`

### ✅ ISSUE 2: Missing error handling
**Fix**: Added proper error handling throughout with logging

### ✅ ISSUE 3: No logging middleware
**Fix**: Added structured logging using standard `log` package

### ✅ ISSUE 4-5: Version parsing bug
**Fix**: Added `strings.TrimSuffix(version, "-SNAPSHOT")` to strip suffix

### ✅ ISSUE 6: No HTTP method validation
**Fix**: Added method check in handlers to only allow GET requests

### ✅ ISSUE 7: Ignoring JSON encode errors
**Fix**: Added error handling for `json.Encoder.Encode()`

### ✅ ISSUE 8: No health check endpoint
**Fix**: Added `/health` endpoint returning `{"status":"UP"}`

### ✅ ISSUE 9: No graceful shutdown
**Fix**: Implemented graceful shutdown with signal handling (SIGINT, SIGTERM)

### ✅ ISSUE 10: Ignoring server errors
**Fix**: Added proper error handling for `server.ListenAndServe()`

---

## 2. Test Fixes (`main_test.go`)

### ✅ Missing test for -SNAPSHOT stripping
**Fix**: Added test cases validating version normalization

### ✅ Missing test for HTTP method validation
**Fix**: Added tests for POST requests returning 405 Method Not Allowed

### ✅ Missing health endpoint tests
**Fix**: Added comprehensive tests for `/health` endpoint

### ✅ Improved test coverage
**Fix**: Added table-driven tests for version parsing

---

## 3. Dockerfile Fixes

### ✅ ISSUE 1: Wrong Go version
**Fix**: Changed from `golang:1.21-alpine` to `golang:1.22-alpine`

### ✅ ISSUE 2-3: Bad layer caching
**Fix**: Restructured to copy `go.mod` first, then download deps, then copy source

### ✅ ISSUE 4: Missing optimization flags
**Fix**: Added `-ldflags="-w -s"` to reduce binary size

### ✅ ISSUE 5: Using full golang runtime image
**Fix**: Changed runtime stage to `alpine:latest` (minimal image)

### ✅ ISSUE 6: Running as root
**Fix**: Created non-root user `appuser` and switched to it

### ✅ ISSUE 7: Copying unnecessary files
**Fix**: Only copy binary and VERSION file to runtime stage

### ✅ ISSUE 8: No HEALTHCHECK
**Fix**: Added Docker HEALTHCHECK directive using `/health` endpoint

### ✅ ISSUE 9: Hardcoded port
**Fix**: Used `ENV PORT=8080` and `${PORT}` variable

### ✅ ISSUE 10: No signal handling
**Fix**: Application now handles signals for graceful shutdown

---

## 4. Makefile Fixes

### ✅ ISSUE 1: Hardcoded credentials
**Fix**: Use environment variables with `?=` for overrides, require secrets via env vars

### ✅ ISSUE 2: No tool validation
**Fix**: Added `check-tools` target to validate required tools

### ✅ ISSUE 3: No coverage reporting
**Fix**: Added `coverage` target with coverage report generation

### ✅ ISSUE 4: Missing lint target
**Fix**: Added `lint` target with golangci-lint (falls back to go vet)

### ✅ ISSUE 5: No version tagging
**Fix**: Docker images now tagged with `latest`, `VERSION`, and `VERSION-SHA`

### ✅ ISSUE 6-7: Hardcoded credentials and single tag
**Fix**: Credentials from env vars, multiple tags pushed

### ✅ ISSUE 8: No CF CLI validation
**Fix**: Added checks for cf CLI installation and login status

### ✅ ISSUE 9: VERSION not normalized
**Fix**: Strip -SNAPSHOT suffix using `sed` before deployment

### ✅ ISSUE 10: No rollback mechanism
**Fix**: Added `rollback` target using `cf rollback`

### ✅ ISSUE 11: Incomplete clean
**Fix**: Clean now removes coverage files, test artifacts, and Docker images

---

## 5. GitHub Actions Fixes (`.github/workflows/ci.yml`)

### ✅ ISSUE 1: Missing environment variables
**Fix**: Added `env` section at workflow level with GO_VERSION, DOCKER_REGISTRY, IMAGE_NAME

### ✅ ISSUE 2: No concurrency control
**Fix**: Added `concurrency` group to prevent parallel deployments

### ✅ ISSUE 3: Jobs without dependencies
**Fix**: Added proper `needs` dependencies: test/lint need build, docker-build needs test+lint, etc.

### ✅ ISSUE 4: Wrong Go version
**Fix**: Changed to Go 1.22 using environment variable

### ✅ ISSUE 5: Test missing dependency
**Fix**: Added `needs: build` to test job

### ✅ ISSUE 6: Missing lint target
**Fix**: Makefile now has lint target

### ✅ ISSUE 7: Docker-build without test dependency
**Fix**: Added `needs: [test, lint]`

### ✅ ISSUE 8: No Docker layer caching
**Fix**: Added Docker Buildx with GitHub Actions cache

### ✅ ISSUE 9: Docker-push missing dependency
**Fix**: Added `needs: docker-build`

### ✅ ISSUE 10: Hardcoded Docker credentials
**Fix**: Use GitHub Secrets via `docker/login-action`

### ✅ ISSUE 11: No proper image tagging
**Fix**: Tag with version (normalized) and commit SHA

### ✅ ISSUE 12: Deploy missing dependency
**Fix**: Added `needs: docker-push`

### ✅ ISSUE 13: Hardcoded CF credentials
**Fix**: Use GitHub Secrets for all CF credentials

### ✅ ISSUE 14: VERSION with -SNAPSHOT
**Fix**: Normalize VERSION before deployment using sed

### ✅ ISSUE 15: No deployment verification
**Fix**: Added smoke tests hitting /health and /info endpoints

### ✅ ISSUE 16: No rollback mechanism
**Fix**: Added rollback step on deployment failure

---

## 6. Cloud Foundry Manifest Fixes (`manifest.yml`)

### ✅ ISSUE 1: VERSION with -SNAPSHOT
**Fix**: Set to normalized version (1.189.0), CI/CD updates it

### ✅ ISSUE 2: No health check configured
**Fix**: Added `health-check-type: http` and `health-check-http-endpoint: /health`

### ✅ ISSUE 3: No route configuration
**Fix**: Added explicit route configuration

### ✅ ISSUE 4: Missing resource limits
**Fix**: Added memory (256M), disk_quota (512M), and process configuration

### ✅ ISSUE 5: No environment-specific config
**Fix**: Added proper resource limits and scaling configuration

---

## Additional Improvements

### Security Enhancements
- Non-root Docker user
- Secrets via environment variables/GitHub Secrets
- No hardcoded credentials anywhere

### Performance Optimizations
- Docker layer caching
- Optimized binary size with build flags
- GitHub Actions caching

### Reliability Improvements
- Graceful shutdown handling
- Health check endpoints
- Deployment verification
- Rollback mechanism
- Proper error handling throughout

### Code Quality
- Comprehensive test coverage
- Linting integration
- Coverage reporting
- Method validation

---

## Required GitHub Secrets

To run the CI/CD pipeline, configure these secrets in GitHub repository settings:

- `DOCKER_REGISTRY` - Docker registry URL (optional, defaults to myregistry.example.com)
- `DOCKER_USERNAME` - Docker registry username
- `DOCKER_PASSWORD` - Docker registry password
- `CF_API` - Cloud Foundry API endpoint
- `CF_USERNAME` - Cloud Foundry username
- `CF_PASSWORD` - Cloud Foundry password
- `CF_ORG` - Cloud Foundry organization
- `CF_SPACE` - Cloud Foundry space

---

## Testing the Fixes

### Local Testing
```bash
# Build and test
make build
make test
make coverage
make lint

# Docker
make docker-build
docker run -p 8080:8080 go-app-demo:latest

# Test endpoints
curl http://localhost:8080/info
curl http://localhost:8080/health
```

### Verify Version Normalization
```bash
# Should return "1.189.0" not "1.189.0-SNAPSHOT"
curl http://localhost:8080/info | jq .version
```

### CI/CD Pipeline
Push to main branch to trigger full pipeline with all fixes applied.

---

**All 52+ issues have been resolved! 🎉**
