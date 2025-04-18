# Build stage
FROM golang:1.22-alpine AS builder

# Add metadata
LABEL maintainer="Your Name <your.email@example.com>"
LABEL description="AWS Secrets Manager Demo Application"

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Final stage
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/main /main

# Copy SSL certificates (needed for AWS SDK)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/secret || exit 1

# Command to run the executable
CMD ["/main"]
