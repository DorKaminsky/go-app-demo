# DevOps Issues - Lite Version

This version contains **10 critical issues** that test core DevOps understanding.

## 📋 Issues Overview

### Go Application (`main.go`) - 6 Issues

**ISSUE 1: Deprecated Package Usage**
- **Line**: 6
- **Problem**: Using deprecated `ioutil.ReadFile` (deprecated since Go 1.16)
- **Impact**: Code quality, future compatibility
- **Fix**: Replace with `os.ReadFile`

**ISSUE 2: Version Not Normalized**
- **Line**: 34
- **Problem**: Not stripping `-SNAPSHOT` suffix from VERSION
- **Impact**: Deployment shows "1.189.0-SNAPSHOT" instead of "1.189.0"
- **Fix**: Add `strings.TrimSuffix(version, "-SNAPSHOT")`

**ISSUE 3: Ignoring JSON Encode Error**
- **Line**: 44
- **Problem**: Not checking if `json.Encoder.Encode()` fails
- **Impact**: Silent failures, no error logging
- **Fix**: Check error and handle appropriately

**ISSUE 4: Missing Health Check Endpoint**
- **Lines**: 47-57
- **Problem**: No `/health` endpoint for Cloud Foundry health checks
- **Impact**: Cloud Foundry can't monitor application health
- **Fix**: Add `healthHandler` and register `/health` route

**ISSUE 5: No Graceful Shutdown**
- **Line**: 60
- **Problem**: Server stops immediately on SIGTERM/SIGINT
- **Impact**: In-flight requests are dropped, data loss possible
- **Fix**: Implement signal handling with `server.Shutdown()`

**ISSUE 6: Ignoring Server Error**
- **Line**: 61
- **Problem**: Not checking `http.ListenAndServe()` error
- **Impact**: Server startup failures go unnoticed
- **Fix**: Check and handle error appropriately

---

### Dockerfile - 4 Issues

**ISSUE 7: Wrong Go Version**
- **Line**: 2
- **Problem**: Using `golang:1.21-alpine` instead of `golang:1.22-alpine`
- **Impact**: Not using required Go version
- **Fix**: Change to `golang:1.22-alpine`

**ISSUE 8: Poor Docker Layer Caching**
- **Lines**: 6-7
- **Problem**: Copying all files before `go mod download`
- **Impact**: Cache invalidated on any code change, slow builds
- **Fix**: Copy `go.mod` first, then download, then copy source
- **Best Practice**:
  ```dockerfile
  COPY go.mod ./
  RUN go mod download
  COPY main.go ./
  ```

**ISSUE 9: Running as Root User**
- **Line**: 19
- **Problem**: Container runs as root (security risk)
- **Impact**: If compromised, attacker has root access
- **Fix**: Create non-root user and switch with `USER appuser`
- **Security Best Practice**: Never run containers as root

**ISSUE 10: No HEALTHCHECK**
- **Line**: 29
- **Problem**: Docker has no way to check container health
- **Impact**: Orchestrators can't detect unhealthy containers
- **Fix**: Add `HEALTHCHECK` directive using `/health` endpoint

---

### CI/CD Pipeline (`.github/workflows/ci.yml`) - 3 Issues

**ISSUE 11: Missing Job Dependency**
- **Line**: 35
- **Problem**: `test` job missing `needs: build`
- **Impact**: Tests may run before build completes
- **Fix**: Add `needs: build` to ensure proper ordering

**ISSUE 12: No Docker Layer Caching**
- **Lines**: 85-91
- **Problem**: Not using GitHub Actions cache for Docker layers
- **Impact**: Every build pulls all layers, very slow
- **Fix**: Add `cache-from: type=gha` and `cache-to: type=gha,mode=max`
- **Performance Impact**: Can reduce build time from 5min to 30sec

**ISSUE 13: Hardcoded Credentials (CRITICAL!)**
- **Line**: 105-107
- **Problem**: Hardcoded password in pipeline
- **Impact**: **CRITICAL SECURITY VULNERABILITY** - credentials exposed in version control
- **Fix**: Use GitHub Secrets: `${{ secrets.DOCKER_PASSWORD }}`
- **Security**: This is the #1 security issue - credentials must NEVER be in code

---

## 🎯 What These Issues Test

### Security Understanding (2 issues)
- ISSUE 9: Docker root user
- ISSUE 13: Hardcoded credentials

### Production Readiness (2 issues)
- ISSUE 4: Health check endpoint
- ISSUE 5: Graceful shutdown

### Performance Optimization (1 issue)
- ISSUE 8: Docker layer caching
- ISSUE 12: CI/CD caching

### Code Quality (2 issues)
- ISSUE 1: Deprecated packages
- ISSUE 2: Version normalization bug

### DevOps Best Practices (3 issues)
- ISSUE 7: Version management
- ISSUE 10: Container health checks
- ISSUE 11: Pipeline dependencies

---

## ✅ Expected Fixes Summary

Good candidates should:
1. ✅ Identify all 10 issues (or at least 8+)
2. ✅ Fix the CRITICAL security issue (#13) first
3. ✅ Implement health check endpoint and graceful shutdown
4. ✅ Optimize Docker builds with proper caching
5. ✅ Demonstrate understanding of why each fix matters

---

## 📊 Difficulty Assessment

**Easy (2-3 issues)**
- ISSUE 1: Deprecated package (obvious)
- ISSUE 7: Wrong Go version (explicit requirement)

**Medium (4-5 issues)**
- ISSUE 2: Version normalization (business logic)
- ISSUE 3, 6: Error handling (code quality)
- ISSUE 11: Job dependencies (CI/CD understanding)

**Hard (2-3 issues)**
- ISSUE 4, 5: Health checks + graceful shutdown (production patterns)
- ISSUE 8, 12: Layer caching (performance optimization)
- ISSUE 9: Non-root user (security best practices)
- ISSUE 13: Hardcoded creds (should be caught immediately!)

---

**Total Time Expected: 1-2 hours** for mid-level engineer with AI assistance
