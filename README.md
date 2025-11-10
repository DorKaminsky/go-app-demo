# Go App Demo - DevOps Engineer Assessment

## 📋 Overview

Welcome to the DevOps Engineer skills assessment! This repository contains a simple Go microservice that exposes a REST API. Your task is to identify and fix DevOps issues in the codebase, containerization, and CI/CD pipeline to make this application production-ready.

## 🎯 The Challenge

This repository contains a Go microservice that **has problems**. Your job is to identify and fix all issues to make this application production-ready.

**Important**: There are issues throughout the codebase. Don't assume anything works correctly just because it runs. Review everything carefully.

## 🏗️ What You're Working With

### Application Details
- **Language**: Go 1.22
- **Service**: Simple REST API microservice
- **Endpoint**: `GET /info` - Returns version and deployment information

**Expected Response:**
```json
{
  "version": "1.189.0",
  "deployed_at": "2024-11-10T10:00:00Z"
}
```

### Project Structure
```
.
├── main.go                    # Main application code
├── main_test.go              # Unit tests
├── go.mod                    # Go module definition
├── VERSION                   # Version file
├── Dockerfile                # Container image definition
├── Makefile                  # Build automation
├── manifest.yml              # Cloud Foundry deployment config
├── .github/workflows/ci.yml  # CI/CD pipeline
└── README.md                 # This file
```

## 🚀 Getting Started

### Prerequisites
- Go 1.22+
- Docker
- Make
- Git

### Local Development
```bash
# Clone the repository
git clone <repository-url>
cd go-app-demo

# Build the application
make build

# Run tests
make test

# Run locally
./go-app-demo

# Test the endpoint
curl http://localhost:8080/info
```

## 🎯 What "Production-Ready" Means

Your fixed version should meet these criteria:

### **Code Quality**
- Follows Go best practices
- Proper error handling throughout
- No deprecated packages
- Appropriate logging
- Graceful shutdown capability
- Health check endpoint available

### **Security**
- **Zero hardcoded credentials** anywhere in the code or configuration
- Secrets managed through environment variables or secure vaults
- Docker containers don't run as root
- No sensitive information in version control

### **Build & Deployment**
- Docker image is optimized (small size, fast builds)
- Version numbers are clean and semantic
- Deployment configuration is complete and correct
- All automation targets work correctly

### **Reliability**
- All tests pass
- Good test coverage
- CI/CD pipeline runs successfully
- Application starts and responds correctly
- Health checks work

### **Best Practices**
- Code is maintainable
- Configuration is environment-agnostic
- Documentation is accurate
- Build process is reproducible

## 📝 Deliverables

### 1. Fixed Code
Create a branch `fix/devops-issues` with all your fixes

### 2. Documentation (FIXES.md)
Create a `FIXES.md` file documenting:
- **Every issue you found** (be thorough - there are many!)
- **How you fixed each issue**
- **Why your solution is better**
- **Any trade-offs or decisions made**
- **Which files you modified and why**

### 3. Working Pipeline
- All CI/CD jobs should pass
- Docker image should build successfully
- Tests should pass with good coverage

### 4. Verification
Demonstrate that:
- Application builds and runs locally
- Application runs via Docker
- Health check endpoint works
- Version is correctly normalized
- No hardcoded secrets remain
- Deployment configuration is production-ready (even if not deployed)

**Note on Deployment**:
- The Cloud Foundry deployment step is commented out in the workflow as it requires SAP internal network access
- However, you should still fix the deployment configuration (`manifest.yml`, Makefile deploy target, etc.)
- The deployment code should be production-ready even if you can't test it live
- You can validate deployment logic locally using `cf push --no-start` or similar dry-run approaches

## 🧪 How to Test Your Fixes

### Local Testing
```bash
# Build and test
make build
make test
make coverage
make lint

# Docker testing
make docker-build
docker run -p 8080:8080 go-app-demo:latest

# Verify endpoints
curl http://localhost:8080/info
curl http://localhost:8080/health

# Check version is normalized (should be "1.189.0", not "1.189.0-SNAPSHOT")
curl http://localhost:8080/info | jq .version
```

### Deployment Configuration Testing
```bash
# Verify manifest.yml is properly configured
cat manifest.yml

# Check Makefile deploy target
make deploy --dry-run || echo "Verify deploy target exists and logic is correct"

# Validate VERSION normalization
cat VERSION
# Should show: 1.189.0 (without -SNAPSHOT after your fixes)
```

### CI/CD Testing
1. Push your changes to your branch
2. Create a pull request
3. Verify all CI/CD jobs pass:
   - ✅ Build
   - ✅ Test (with coverage)
   - ✅ Lint
   - ✅ Docker Build
   - ✅ Docker Push (on main branch)
   - 📝 Deploy step (commented out, but code should be correct)

## 📚 Evaluation Criteria

Your submission will be evaluated on:

1. **Completeness** (40%)
   - How many issues did you identify and fix?
   - Are all required fixes implemented?

2. **Code Quality** (25%)
   - Are your fixes following best practices?
   - Is the code clean and maintainable?
   - Proper error handling and logging?

3. **Security** (20%)
   - Are all hardcoded secrets removed?
   - Proper security practices implemented?
   - Following principle of least privilege?

4. **Documentation** (10%)
   - Is your FIXES.md clear and comprehensive?
   - Did you explain your decisions?

5. **Bonus Features** (5%)
   - Did you implement extra improvements?
   - Creative solutions to problems?

## 🕐 Time Expectation

This assessment should take approximately **3-4 hours** for an experienced DevOps engineer.

**Tip**: Take your time to review every file thoroughly. The issues are spread across multiple files and areas.

## ❓ Getting Help

If you have questions:
1. Check the documentation in the codebase
2. Review error messages carefully
3. Google is your friend (just like in real work!)
4. If stuck, document what you've tried

## 📤 Submission

When you're ready to submit:

1. Ensure all changes are committed to your `fix/devops-issues` branch
2. Create `FIXES.md` documenting all your changes
3. Create a pull request to main branch
4. In the PR description, include:
   - Summary of issues found and fixed
   - Any challenges you faced
   - Time spent on the assessment
   - Explanation of deployment configuration fixes (even if not live-tested)

## 🎯 Success Criteria

A successful submission should have:
- ✅ All critical issues fixed
- ✅ All CI/CD jobs passing (except deploy which is disabled)
- ✅ No hardcoded secrets
- ✅ Docker image builds and runs correctly
- ✅ Application starts and responds to requests
- ✅ Health check endpoint working
- ✅ Version properly normalized
- ✅ Deployment configuration is production-ready
- ✅ Comprehensive documentation in FIXES.md

---

**Good luck! We're excited to see your DevOps skills in action! 🚀**

## 💡 Where to Start

Not sure where to begin? Here's a suggested approach:

1. **Run everything first** - Try building, testing, and running the application. What fails? What warnings do you see?

2. **Read the code carefully** - Look at every file. Are there any comments? Deprecated warnings? Security issues?

3. **Test the CI/CD pipeline** - Push to a branch and see what happens. Does the pipeline make logical sense?

4. **Check the configuration files** - Are credentials handled securely? Are versions correct?

5. **Compare with best practices** - How does this codebase compare to production-ready Go applications you've seen?

Remember: **There are issues in every file.** Don't stop after finding a few problems!

---

**Questions?** Feel free to reach out to the hiring team.
