# Go App Demo - DevOps Candidate Assessment

## 🎯 Overview

This is a DevOps skills assessment project. The repository contains a simple Go microservice with **intentional DevOps issues** that need to be identified and fixed.

The service exposes a single REST API endpoint that returns version and deployment information.

## 📋 Service Specification

- **Service Name**: `go-app-demo`
- **Language**: Go 1.22
- **Endpoint**: `GET /info`
- **Response Format**:
```json
{
  "version": "1.189.0",
  "deployed_at": "2024-11-05T10:00:00Z"
}
```

## 🏗️ Project Structure

```
.
├── main.go                    # Main application code
├── main_test.go              # Unit tests
├── go.mod                    # Go module definition
├── VERSION                   # Version file (1.189.0-SNAPSHOT)
├── Dockerfile                # Container image definition
├── Makefile                  # Build automation
├── manifest.yml              # Cloud Foundry deployment config
├── .github/workflows/ci.yml  # CI/CD pipeline
└── README.md                 # This file
```

## 🚀 Local Development

### Prerequisites
- Go 1.22+
- Docker
- Make
- Cloud Foundry CLI (for deployment)

### Running Locally
```bash
# Build the application
make build

# Run tests
make test

# Run the application
./go-app-demo

# Test the endpoint
curl http://localhost:8080/info
```

### Docker
```bash
# Build Docker image
make docker-build

# Run container
docker run -p 8080:8080 go-app-demo:latest
```

## 🔧 Deployment

Deploy to Cloud Foundry:
```bash
make deploy
```

## 🐛 INTENTIONAL ISSUES (FOR REVIEWER - DELETE BEFORE GIVING TO CANDIDATE)

### 1. Go Code Issues (`main.go`)
- ❌ **ISSUE 1**: Using deprecated `ioutil` package (should use `os.ReadFile`)
- ❌ **ISSUE 2**: Missing error handling in multiple places
- ❌ **ISSUE 3**: No logging middleware or structured logging
- ❌ **ISSUE 4**: Version parsing bug - doesn't strip `-SNAPSHOT` suffix from VERSION
- ❌ **ISSUE 5**: `getVersion()` returns `1.189.0-SNAPSHOT` instead of `1.189.0`
- ❌ **ISSUE 6**: `infoHandler` doesn't check HTTP method (should only allow GET)
- ❌ **ISSUE 7**: Ignoring error from `json.Encode()`
- ❌ **ISSUE 8**: No `/health` endpoint for Cloud Foundry health checks
- ❌ **ISSUE 9**: No graceful shutdown handling
- ❌ **ISSUE 10**: Ignoring error from `http.ListenAndServe()`

### 2. Test Issues (`main_test.go`)
- ❌ Test expects buggy behavior (doesn't validate -SNAPSHOT stripping)
- ❌ Missing test coverage for version normalization
- ❌ No test for HTTP method validation

### 3. Dockerfile Issues
- ❌ **ISSUE 1**: Using Go 1.21 instead of 1.22
- ❌ **ISSUE 2**: Bad layer caching - copying all files before `go mod download`
- ❌ **ISSUE 3**: Should copy `go.mod` first, then download deps, then copy code
- ❌ **ISSUE 4**: Building without optimization flags (`-ldflags="-w -s"`)
- ❌ **ISSUE 5**: Using full `golang` image for runtime instead of minimal `alpine`
- ❌ **ISSUE 6**: Running as root user (security issue)
- ❌ **ISSUE 7**: Copying unnecessary files
- ❌ **ISSUE 8**: No `HEALTHCHECK` defined
- ❌ **ISSUE 9**: Hardcoded port instead of using `ENV`
- ❌ **ISSUE 10**: No signal handling for graceful shutdown

### 4. Makefile Issues
- ❌ **ISSUE 1**: Hardcoded registry URL and credentials
- ❌ **ISSUE 2**: No validation of required tools (docker, cf, go)
- ❌ **ISSUE 3**: No coverage report or coverage threshold
- ❌ **ISSUE 4**: Missing `lint` target that CI expects
- ❌ **ISSUE 5**: Docker build only uses `latest` tag, no version tagging
- ❌ **ISSUE 6**: Hardcoded credentials in `docker-push` (SECURITY!)
- ❌ **ISSUE 7**: Only pushing `latest` tag, not version-specific tag
- ❌ **ISSUE 8**: No check if CF CLI is installed
- ❌ **ISSUE 9**: VERSION not normalized (still has -SNAPSHOT)
- ❌ **ISSUE 10**: No rollback mechanism
- ❌ **ISSUE 11**: `clean` target incomplete

### 5. GitHub Actions Issues (`.github/workflows/ci.yml`)
- ❌ **ISSUE 1**: Missing environment variables at workflow level
- ❌ **ISSUE 2**: No concurrency control
- ❌ **ISSUE 3**: Jobs run in parallel without dependencies (race conditions)
- ❌ **ISSUE 4**: Using Go 1.21 instead of 1.22
- ❌ **ISSUE 5**: `test` job missing `needs: build`
- ❌ **ISSUE 6**: `make lint` target doesn't exist in Makefile
- ❌ **ISSUE 7**: `docker-build` should depend on tests passing
- ❌ **ISSUE 8**: No Docker layer caching configured
- ❌ **ISSUE 9**: `docker-push` missing dependency on `docker-build`
- ❌ **ISSUE 10**: Hardcoded Docker credentials (SECURITY!)
- ❌ **ISSUE 11**: No proper image tagging (version/SHA)
- ❌ **ISSUE 12**: `deploy` missing dependency on `docker-push`
- ❌ **ISSUE 13**: Hardcoded CF credentials (SECURITY!)
- ❌ **ISSUE 14**: VERSION has -SNAPSHOT, wrong version deployed
- ❌ **ISSUE 15**: No deployment verification/smoke tests
- ❌ **ISSUE 16**: No rollback mechanism

### 6. Cloud Foundry Manifest Issues (`manifest.yml`)
- ❌ **ISSUE 1**: VERSION env var still has `-SNAPSHOT`
- ❌ **ISSUE 2**: No health check endpoint configured
- ❌ **ISSUE 3**: No route configuration
- ❌ **ISSUE 4**: Missing resource limits
- ❌ **ISSUE 5**: No environment-specific config

## 🎓 Candidate Tasks

**Your mission**: Fix the DevOps issues in this repository to make it production-ready.

### Required Fixes:
1. ✅ Fix version normalization (strip `-SNAPSHOT` suffix)
2. ✅ Fix Dockerfile layer caching for faster builds
3. ✅ Remove all hardcoded secrets and use GitHub Secrets
4. ✅ Fix CI/CD pipeline job dependencies
5. ✅ Add proper Docker image tagging (version + SHA)
6. ✅ Add health check endpoint and configure it in manifest
7. ✅ Fix Go code issues (error handling, deprecated packages)
8. ✅ Add linting to Makefile and fix lint errors
9. ✅ Implement graceful shutdown
10. ✅ Add deployment verification

### Bonus Points:
- Add rollback mechanism
- Implement proper logging
- Add code coverage reporting
- Optimize Docker image size
- Add security scanning
- Implement blue-green deployment

## 📝 Submission Guidelines

1. Fork this repository
2. Create a branch: `fix/devops-issues`
3. Fix the issues
4. Document your changes in `FIXES.md`
5. Submit a pull request

## 🔐 Required GitHub Secrets

For CI/CD to work, configure these secrets:
- `DOCKER_REGISTRY` - Docker registry URL
- `DOCKER_USERNAME` - Docker registry username
- `DOCKER_PASSWORD` - Docker registry password
- `CF_API` - Cloud Foundry API endpoint
- `CF_USERNAME` - Cloud Foundry username
- `CF_PASSWORD` - Cloud Foundry password
- `CF_ORG` - Cloud Foundry organization
- `CF_SPACE` - Cloud Foundry space

## 📚 Resources

- [Go Best Practices](https://golang.org/doc/effective_go)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Cloud Foundry Documentation](https://docs.cloudfoundry.org/)

---

**Good luck! 🚀**