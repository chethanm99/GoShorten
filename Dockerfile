# -------------------------
# Stage 1: Builder
# -------------------------
FROM golang:alpine AS builder

# Set working directory inside the container
WORKDIR /src

# Copy go.mod and go.sum first (to leverage Docker layer caching)
COPY app/go.mod app/go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY app/ .

# Build the Go binary (output to /build/main)
RUN go build -o /build/main .

# -------------------------
# Stage 2: Final image
# -------------------------
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S -G appgroup appuser

# Copy the compiled binary from builder stage
COPY --from=builder --chown=appuser:appgroup /build/main /app/main

# Set working directory
WORKDIR /app

# Switch to the non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://localhost:8080/health || exit 1

# Start the application
CMD ["/app/main"]
