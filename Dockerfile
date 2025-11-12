# Build stage
FROM golang:1.22-alpine AS builder

# Copy go mod files first for better layer caching
COPY go.mod ./
RUN go mod download

# Copy source code
COPY main.go ./
COPY VERSION ./

# Build application with optimization flags
RUN go build -ldflags="-w -s" -o go-app-demo .

# Runtime stage
FROM alpine:latest

# Create non-root user for security
RUN adduser -D -u 1000 appuser

# Copy binary and VERSION file from builder
COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port
ENV PORT=8080
EXPOSE ${PORT}

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

# Run application
CMD ["./go-app-demo"]
