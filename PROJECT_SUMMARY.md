# DevOps Candidate Test - Project Summary

## 📁 Complete File Structure

```
go-app-demo/
├── .github/
│   └── workflows/
│       └── ci.yml              # CI/CD pipeline with 16 intentional issues
├── .gitignore                  # Go project gitignore
├── Dockerfile                  # Multi-stage build with 10 intentional issues
├── Makefile                    # Build automation with 11 intentional issues
├── README.md                   # Complete documentation with all issues listed
├── VERSION                     # Contains "1.189.0-SNAPSHOT"
├── go.mod                      # Go module definition
├── main.go                     # Main application with 10 intentional issues
├── main_test.go                # Unit tests with missing coverage
├── manifest.yml                # Cloud Foundry config with 5 intentional issues
└── PROJECT_SUMMARY.md          # This file
```

## 🎯 Quick Stats

- **Total Files Created**: 10
- **Total Intentional Issues**: 52+
- **Languages**: Go, YAML, Makefile, Dockerfile
- **Lines of Code**: ~300+

## 🐛 Issue Categories

### Security Issues (Critical)
- Hardcoded Docker credentials in Makefile
- Hardcoded Docker credentials in GitHub Actions
- Hardcoded Cloud Foundry credentials in GitHub Actions
- Running Docker container as root
- No secrets management

### CI/CD Issues
- Wrong job dependencies (race conditions)
- Missing concurrency control
- No deployment verification
- No rollback mechanism
- Wrong Go version (1.21 vs 1.22)
- Missing lint target

### Docker Issues
- Poor layer caching
- Wrong base image version
- No health check
- Inefficient image size
- Missing optimization flags

### Code Quality Issues
- Deprecated packages (ioutil)
- Missing error handling
- No structured logging
- No graceful shutdown
- Version parsing bug (doesn't strip -SNAPSHOT)

### Cloud Foundry Issues
- VERSION not normalized
- No health check configured
- Missing route configuration
- No environment-specific config

## 🧪 Testing the Repository

### Verify Issues Exist

1. **Version Bug Test**:
```bash
cd /Users/I572966/git/github.tools/cfs-devops/go-app-demo
go run main.go &
curl http://localhost:8080/info
# Should show "1.189.0-SNAPSHOT" instead of "1.189.0"
```

2. **Build Test**:
```bash
make build
# Should succeed
```

3. **Test Run**:
```bash
make test
# Should pass (but tests don't validate -SNAPSHOT stripping)
```

4. **Lint Test**:
```bash
make lint
# Should FAIL - target doesn't exist
```

5. **Docker Build Test**:
```bash
make docker-build
# Should work but be slow due to poor caching
```

## 📋 Candidate Evaluation Criteria

### Must Fix (Core DevOps Skills)
- [ ] Version normalization (strip -SNAPSHOT)
- [ ] Fix Dockerfile layer caching
- [ ] Remove hardcoded secrets
- [ ] Fix CI/CD job dependencies
- [ ] Add proper Docker image tagging

### Should Fix (Intermediate)
- [ ] Add health check endpoint
- [ ] Fix error handling in Go code
- [ ] Add lint target and fix issues
- [ ] Configure CF health checks
- [ ] Add deployment verification

### Nice to Have (Advanced)
- [ ] Implement graceful shutdown
- [ ] Add rollback mechanism
- [ ] Optimize Docker image size
- [ ] Add code coverage reporting
- [ ] Implement structured logging
- [ ] Add security scanning

## 🚀 Next Steps for Reviewer

1. **Test the repository**:
   - Clone and run locally
   - Verify all intentional issues are present
   - Test that fixes would actually work

2. **Before giving to candidate**:
   - Delete the "INTENTIONAL ISSUES" section from README.md
   - Delete this PROJECT_SUMMARY.md file
   - Optionally create a private answer key

3. **Candidate Instructions**:
   - Give them 2-4 hours
   - Ask them to document all issues found
   - Request fixes with explanations
   - Evaluate both identification and solutions

## 📝 Answer Key Location

All issues are documented in README.md under:
**"🐛 INTENTIONAL ISSUES (FOR REVIEWER - DELETE BEFORE GIVING TO CANDIDATE)"**

Delete this section before distributing to candidates!

---

**Repository is ready for testing and deployment! 🎉**
