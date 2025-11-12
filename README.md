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
manifest.yml              # Deployment configuration
VERSION                   # Version file
```

---

## 🚀 Getting Started

### Prerequisites

- [ ] GitHub account
- [ ] Git installed (`git --version`)
- [ ] Go 1.22+ (`go version`)
- [ ] Docker (`docker --version`)
- [ ] Make (`make --version`)
- [ ] Docker Hub account (free at https://hub.docker.com)
- [ ] 3-4 hours of uninterrupted time

### Step 1: Access Your Repository

You've been granted **write access** to a repository created specifically for you:

```
https://github.com/DorKaminsky/go-app-demo-[YOUR-NAME]
```

The hiring team will provide you with the exact repository URL.

### Step 2: Clone and Setup

```bash
# Clone your assigned repository
git clone https://github.com/DorKaminsky/go-app-demo-[YOUR-NAME].git
cd go-app-demo-[YOUR-NAME]

# Create working branch
git checkout -b fix/devops-issues

# Verify environment (expect failures - that's the point!)
make build
make test
```

### Step 3: Set Up GitHub Secrets

For CI/CD to work, add these secrets in your repository:

1. Go to **Settings → Secrets and variables → Actions**
2. Add **New repository secret**:

| Secret Name | Value | How to Get |
|-------------|-------|------------|
| `DOCKERHUB_USERNAME` | Your Docker Hub username | Sign up at hub.docker.com |
| `DOCKERHUB_TOKEN` | Docker Hub access token | Account Settings → Security → New Access Token |

---

## 🎯 Your Mission

### Phase 1: Fix All Issues

Identify and fix bugs in:
- ✅ Go code (deprecated packages, missing features, error handling)
- ✅ Tests (failing assertions, incorrect validations)
- ✅ Dockerfile (missing directives, security issues)
- ✅ Makefile (syntax errors, wrong commands)
- ✅ CI/CD pipeline (deprecated actions, wrong runners, typos)
- ✅ Deployment config (wrong deployment method)

**Production-Ready Criteria:**
- All tests pass
- No hardcoded credentials
- Version normalized (no `-SNAPSHOT`)
- Health endpoint works
- Non-root Docker user
- CI/CD pipeline passes
- Good test coverage (>80%)

### Phase 2: Kubernetes Deployment

After fixing all issues:

1. **Create a Helm chart** in `helm/go-app-demo/` to deploy the application
2. **Create GitHub Actions workflow** (`.github/workflows/deploy.yml`) that deploys to Kubernetes using Helm
3. Deployment must use the Docker image from Phase 1

---

## 🧪 Testing Your Work

```bash
# Local testing
make build
make test
make coverage
make lint

# Docker testing
make docker-build
docker run -p 8080:8080 go-app-demo:latest

# Test endpoints
curl http://localhost:8080/info
curl http://localhost:8080/health

# Verify version is normalized (no -SNAPSHOT)
curl http://localhost:8080/info | jq .version
```

---

## 📝 Submission

### Final Checklist

**Phase 1: Fix Existing Issues**
- [ ] All critical issues fixed
- [ ] Tests pass locally (`make test`)
- [ ] Docker builds and runs (`make docker-build`)
- [ ] No hardcoded secrets
- [ ] Version normalized (no `-SNAPSHOT`)
- [ ] Health endpoint works
- [ ] CI/CD pipeline passes (build, test, lint, docker-build, docker-push)

**Phase 2: Kubernetes Deployment**
- [ ] Helm chart created in `helm/go-app-demo/`
- [ ] GitHub Actions deploy workflow created (`.github/workflows/deploy.yml`)
- [ ] Deploy workflow uses Docker image from Docker Hub

**Final Submission**
- [ ] Create Pull Request from `fix/devops-issues` → `main`
- [ ] PR description includes all fixes and Kubernetes deployment
- [ ] All CI/CD jobs pass (including deploy workflow)

### Create Your Pull Request

```bash
# Commit and push your changes
git add .
git commit -m "fix: resolve all DevOps issues and add Kubernetes deployment"
git push origin fix/devops-issues

# Go to GitHub and create Pull Request:
# From: fix/devops-issues
# To: main
#
# Include in PR description:
# - List of all issues found and fixed
# - Screenshots of successful deployment
# - Evidence of passing CI/CD pipeline
```

**Remember**: Issues exist in MANY files. Don't stop early!

---

## ❓ Troubleshooting

### Docker Hub Rate Limits

If you see "toomanyrequests: You have reached your pull rate limit":

```bash
# Solution: Login to Docker Hub
docker login
# Enter your Docker Hub credentials
```

### Tests Failing

That's expected! Finding and fixing failing tests is part of the assessment.

### CI/CD Not Working

1. Check GitHub Secrets are configured correctly
2. Make sure tests pass locally first
3. Review the error logs in GitHub Actions tab

### Need Help?

- ✅ Use AI tools (ChatGPT, Claude, Copilot)
- ✅ Google for solutions
- ✅ Read error messages carefully
- 📧 Email hiring team for clarification (not answers!)

---

**Good luck! 🚀**
