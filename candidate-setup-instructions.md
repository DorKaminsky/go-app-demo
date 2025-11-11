# Setup Instructions - Go App Demo Assessment

## Welcome! 👋

Thank you for your interest in our DevOps Engineer position! This document will guide you through setting up your environment to complete the technical assessment.

---

## 🎯 Assessment Approach: Fork Workflow

**Approach:** Fork the repository to your own GitHub account to complete this assessment.

**Why Fork?**
- ✅ You have full control over your fork
- ✅ You can commit and push freely without restrictions
- ✅ Easy to share your work via pull request
- ✅ Allows GitHub Actions to run in your account

---

## 📋 Getting Started

### Prerequisites

Before you begin, ensure you have:

- [ ] GitHub account (if you don't have one, create at https://github.com/signup)
- [ ] Git installed locally (`git --version` to check)
- [ ] Go 1.22+ installed (`go version` to check)
- [ ] Docker installed (`docker --version` to check)
- [ ] Make installed (`make --version` to check)
- [ ] A text editor or IDE (VS Code, GoLand, etc.)
- [ ] 3-4 hours of uninterrupted time

---

## 🚀 Step-by-Step Setup

### Step 1: Fork the Repository

1. **Navigate to the repository:**
   - You should have received a link to the repository from the hiring team
   - Example: `https://github.com/ORGANIZATION/go-app-demo`

2. **Click the Fork button:**
   - Look for the "Fork" button in the top-right corner of the page
   - Click it

3. **Configure your fork:**
   - **Owner:** Select your personal GitHub account
   - **Repository name:** Keep it as `go-app-demo` (or rename if you prefer)
   - **Description:** (Optional)
   - ⚠️ **IMPORTANT:** check "Copy the main branch only"
   - Click "Create fork"

4. **Wait for fork to complete:**
   - GitHub will create your personal copy
   - You'll be redirected to `https://github.com/[YOUR-USERNAME]/go-app-demo`

---

### Step 2: Clone Your Fork

Open your terminal and run:

```bash
# Clone your fork (replace [YOUR-USERNAME] with your GitHub username)
git clone https://github.com/[YOUR-USERNAME]/go-app-demo.git

# Navigate into the directory
cd go-app-demo

# Verify you're on the main branch
git branch
# Should show: * main
```

**Verify your remote:**
```bash
git remote -v
# Should show your fork as origin
```

---

### Step 3: Create Your Working Branch

Create a new branch for your fixes:

```bash
# Create and switch to your working branch
git checkout -b fix/devops-issues

# Verify you're on the new branch
git branch
# Should show: * fix/devops-issues
```

---

### Step 4: Read the README

Before starting, thoroughly read the README:

```bash
# Read the README
cat README.md

# Or open it in your editor
code README.md  # VS Code
# or
open README.md  # macOS
```

**Key sections to focus on:**
- What "Production-Ready" means
- Deliverables expected
- Evaluation criteria
- Testing instructions

---

### Step 5: Verify Your Environment

Test that everything works:

```bash
# Try building (it might fail - that's expected!)
make build

# Try running tests (some might fail - find out why!)
make test

# Check if Docker works
docker --version
docker ps
```

**Note:** If things fail at this stage, that's normal! Finding and fixing issues is the whole point of this assessment.

---

- ✅ Use AI tools (ChatGPT, Claude, Copilot, etc.) if you want
- ✅ Google for best practices and solutions
- ✅ Test your changes thoroughly
- ✅ Commit frequently with clear messages
- ✅ Ask questions if you need clarification
---

## 💾 Committing Your Changes

As you fix issues, commit your changes:

```bash
# Check what you've changed
git status

# Add files to commit
git add main.go
git add Dockerfile
# Or add all: git add .

# Commit with a clear message
git commit -m "fix: replace deprecated ioutil with os.ReadFile"

# Push to your fork
git push origin fix/devops-issues
```

**Tip:** Make multiple commits! We want to see your thought process.

Good commit messages:
- ✅ `fix: remove hardcoded credentials from CI pipeline`
- ✅ `feat: add health check endpoint`
- ✅ `refactor: implement graceful shutdown`

Bad commit messages:
- ❌ `fixed stuff`
- ❌ `updates`
- ❌ `asdf`

---

## 🔐 Setting Up GitHub Secrets (For CI/CD)

To get the CI/CD pipeline working in your fork, you'll need to configure some secrets:

1. **Go to your fork on GitHub**
2. **Click Settings → Secrets and variables → Actions**
3. **Add the following secrets:**

**Required Secrets:**

| Secret Name | Description | Example Value |
|-------------|-------------|---------------|
| `DOCKERHUB_USERNAME` | Your Docker Hub username | `your-dockerhub-username` |
| `DOCKERHUB_TOKEN` | Docker Hub access token | (Create at hub.docker.com) |

### Creating a Docker Hub Token:

1. Go to https://hub.docker.com
2. Sign up/login
3. Go to Account Settings → Security → New Access Token
4. Name: `github-actions`
5. Copy the token and add it to GitHub Secrets

---

## ❓ Troubleshooting

### Issue: "Can't fork the repository"

**Solution:**
- Make sure you have access to the original repository
- Try refreshing the page
- Make sure you're logged into GitHub
- Contact the hiring team if the issue persists

### Issue: "Tests are failing"

**That's expected!** Finding and fixing failing tests is part of the assessment.

**Debugging steps:**
1. Read the error message carefully
2. Check what the test expects vs what the code does
3. Fix the underlying issue
4. Re-run tests

### Issue: "Docker build is slow"

**Tips:**
- Check your Docker layer caching
- First build is always slow (downloading base images)
- Subsequent builds should be faster

### Issue: "CI/CD pipeline is failing"

**Common causes:**
1. Missing GitHub Secrets (see "Setting Up GitHub Secrets" section)
2. Docker Hub rate limits (see below)
3. Test failures (fix the tests first)
4. Build errors (make sure it works locally first)

### Issue: "Docker Hub rate limits"

**Error:** "toomanyrequests: You have reached your pull rate limit"

**Solution:**
```bash
# Login to Docker Hub (creates/uses free account)
docker login

# Enter your Docker Hub username and password
# This significantly increases your rate limit
```

**For CI/CD**: Make sure you've set up `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` in GitHub Secrets.

---

## 🎉 You're All Set!