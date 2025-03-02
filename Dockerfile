# Use the official Golang 1.18 image as the base image
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Use a minimal base image for the final container
FROM debian:buster-slim

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose application port (change if necessary)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
