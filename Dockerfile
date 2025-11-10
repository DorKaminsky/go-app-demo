# Build stage
FROM golang:1.21-alpine AS builder  # ISSUE 7: Wrong Go version (should be 1.22)

WORKDIR /app

# ISSUE 8: Bad layer caching - copying everything before go mod download
COPY . .

RUN go mod download

# Build application
RUN go build -o go-app-demo .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# ISSUE 9: Running as root user (security risk)
# Should create non-root user

# Copy binary and VERSION
COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

ENV PORT=8080
EXPOSE ${PORT}

# ISSUE 10: No HEALTHCHECK defined
# Docker has no way to check if container is healthy

CMD ["./go-app-demo"]
