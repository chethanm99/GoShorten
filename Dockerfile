# -------------------------
# Stage 1: Builder
# -------------------------
FROM golang:alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker layer caching)
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary (output to /build/main)
RUN go build -o /build/main .

# -------------------------
# Stage 2: Final image
# -------------------------
FROM alpine

# Create a non-root user for security
RUN adduser -S -D -H appuser

# Switch to the non-root user
USER appuser

# Copy the compiled binary from builder stage
COPY --from=builder /build/main /app/main

# Set working directory
WORKDIR /app

# Copy environment file if your app needs it
COPY .env .env

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["/app/main"]
