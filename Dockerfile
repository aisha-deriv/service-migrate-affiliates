# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o bin/service-migrate-affiliates ./cmd/service-migrate-affiliates

# Stage 2: Create the final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/bin/service-migrate-affiliates .

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./service-migrate-affiliates"]
