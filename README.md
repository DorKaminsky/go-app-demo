# Go App Demo - DevOps Engineer Assessment

## 📋 Overview

This is a Go microservice with intentional DevOps issues. Your task: identify and fix all issues to make it production-ready.

**Important**: Issues exist in every file. Review everything carefully.

## 🏗️ What You're Working With

**Application**: Simple REST API microservice (Go 1.22)
- `GET /info` - Returns version and deployment information
- `GET /health` - Health check endpoint (after you add it!)

**Expected Response:**
```json
{
  "version": "1.189.0",
  "deployed_at": "2024-11-10T10:00:00Z"
}
```

**Files to Review:**
```
main.go                    # Application code
main_test.go              # Unit tests
Dockerfile                # Container configuration
Makefile                  # Build automation
.github/workflows/ci.yml  # CI/CD pipeline
manifest.yml              # Cloud Foundry deployment
VERSION                   # Version file
```

## 🚀 Getting Started

**Prerequisites**: Go 1.22+, Docker, Make, Git

**Setup:**
```bash
# Fork the repo first! See candidate-setup-instructions.md

# Clone your fork
git clone https://github.com/[YOUR-USERNAME]/go-app-demo.git
cd go-app-demo

# Create working branch
git checkout -b fix/devops-issues

# Try running (expect failures!)
make build
make test
./go-app-demo
curl http://localhost:8080/info
```

## 🎯 Production-Ready Criteria

### Code Quality
- Go best practices, no deprecated packages
- Proper error handling and logging
- Graceful shutdown
- Health check endpoint

### Security
- **Zero hardcoded credentials**
- Secrets via environment variables
- Non-root Docker user
- No sensitive data in git

### Build & Deployment
- Optimized Docker image
- Clean version numbers (no `-SNAPSHOT`)
- Working CI/CD pipeline
- Valid deployment configuration

### Testing
- All tests pass
- Good coverage (above 80%)
- Linter passes

## 📝 Deliverables
1. **Working Pipeline**: All CI/CD jobs pass (except deploy)
2. **Pull Request**: To your fork's main branch with summary

## 🧪 Testing Your Fixes

```bash
# Local testing
make build
make test
make coverage
make lint

# Docker testing
make docker-build
docker run -p 8080:8080 go-app-demo:latest
curl http://localhost:8080/info
curl http://localhost:8080/health

# Verify version normalization
curl http://localhost:8080/info | jq .version
# Should return: "1.189.0" (not "1.189.0-SNAPSHOT")
```

## 🔧 About Cloud Foundry Deployment

**Note**: The Cloud Foundry deployment in the CI/CD pipeline is commented out because it requires internal network access.

**However**, you should still fix the deployment configuration:

### What to Fix:
1. **manifest.yml**:

2. **Makefile `deploy` target**:


### How to Validate (Without Actually Deploying):
```bash
# Check Makefile deploy logic
make --dry-run deploy
# Should show VERSION normalization

# Validate manifest.yml syntax
cat manifest.yml
# Check: health-check-http-endpoint: /health
# Check: VERSION without -SNAPSHOT
```

**Why This Matters**: Shows you understand Cloud Foundry deployment even if you can't test it live.

## ✅ Success Checklist

Before submitting:
- [ ] All critical issues fixed
- [ ] Tests pass locally (`make test`)
- [ ] Docker builds and runs (`make docker-build`)
- [ ] No hardcoded secrets
- [ ] Version normalized (no `-SNAPSHOT`)
- [ ] Health endpoint works
- [ ] CI/CD pipeline passes (except deploy job)
- [ ] Deployment config is production-ready


**Remember**: Issues exist in MANY file. Don't stop early!

## ❓ Help

- ✅ Use AI tools (ChatGPT, Claude, Copilot)
- ✅ Google for solutions
- ✅ Read error messages carefully
- 📧 Email hiring team for clarification (not answers!)

---

**Good luck! 🚀**
