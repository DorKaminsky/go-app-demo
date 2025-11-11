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
1. **Working Pipeline**: All CI/CD jobs pass (except deploy - see below)
2. **Pull Request**: Create a PR from `fix/devops-issues` → `main` **in your fork** (NOT to the original repository)
   - Clear summary of all fixes
   - Description of issues found and resolved

**Note**: To get CI/CD working, you'll need to set up GitHub Secrets. See the **"Setting Up GitHub Secrets"** section in `candidate-setup-instructions.md`.

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

**Note**: This assignment does NOT require a working Cloud Foundry deployment job in the CI/CD pipeline. You won't need actual Cloud Foundry access.

**What You Need to Fix**:
1. **manifest.yml** - Ensure the configuration is production-ready:
   - Health check endpoint points to `/health`
   - VERSION is normalized (no `-SNAPSHOT`)
   - All configuration is valid

2. **Makefile `deploy` target** - Verify the deployment logic:
   - VERSION normalization works correctly
   - Commands are properly structured

### How to Validate (Without Actually Deploying):
```bash
# Check Makefile deploy logic (dry-run)
make --dry-run deploy
# Should show VERSION normalization: 1.189.0 (not 1.189.0-SNAPSHOT)

# Validate manifest.yml syntax
cat manifest.yml
# Verify: health-check-http-endpoint: /health
# Verify: VERSION is normalized
# Verify: All YAML is valid
```

**Why This Matters**: Demonstrates you understand Cloud Foundry deployment configuration and can prepare deployment-ready manifests, even without live access to a CF environment.

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

## ❓ Help & Troubleshooting

### Getting Help
- ✅ Use AI tools (ChatGPT, Claude, Copilot)
- ✅ Google for solutions
- ✅ Read error messages carefully
- 📧 Email hiring team for clarification (not answers!)

### Common Issues

**Docker Hub Rate Limits**

If you encounter errors like "toomanyrequests: You have reached your pull rate limit":

```bash
# Solution: Login to Docker Hub (free account)
docker login

# Enter your Docker Hub username and password
# This increases your rate limit significantly
```

**Alternative**: Use authenticated pulls in your Dockerfile or switch to a different base image registry.

---

**Good luck! 🚀**
