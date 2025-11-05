# ISSUE 1: Using older Go version instead of 1.22
FROM golang:1.21-alpine AS builder

WORKDIR /app

# ISSUE 2: Bad layer caching - copying all files before downloading dependencies
# This means any code change invalidates the dependency cache
COPY . .

# ISSUE 3: Should copy go.mod and go.sum first, then run go mod download
RUN go mod download

# ISSUE 4: Building without optimization flags
RUN go build -o go-app-demo .

# ISSUE 5: Using full golang image for runtime instead of minimal alpine
FROM golang:1.21-alpine

WORKDIR /app

# ISSUE 6: Running as root user - security issue
# Should create non-root user

# ISSUE 7: Copying unnecessary files
COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

# ISSUE 8: No health check defined
# HEALTHCHECK should be added

# ISSUE 9: Hardcoded port instead of using ENV
EXPOSE 8080

# ISSUE 10: No signal handling for graceful shutdown
CMD ["./go-app-demo"]
