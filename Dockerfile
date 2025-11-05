# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod first for better layer caching
COPY go.mod ./

# Download dependencies (cached if go.mod hasn't changed)
RUN go mod download

# Copy source code
COPY main.go ./
COPY VERSION ./

# Build with optimization flags to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o go-app-demo .

# Runtime stage - use minimal alpine image
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Copy only the binary and VERSION file
COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Use environment variable for port
ENV PORT=8080
EXPOSE ${PORT}

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

# Application handles graceful shutdown via signals
CMD ["./go-app-demo"]
