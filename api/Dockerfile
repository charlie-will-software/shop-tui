# Multi-stage Dockerfile for a Go application

# Stage 1: Build stage
FROM golang:latest AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app main.go

# Stage 2: Runtime stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/app .

# Expose the application port (adjust as necessary)
EXPOSE 8080

# Copy environment variables from a .env file if available
# Uncomment the following line if you plan to use a .env file
# COPY .env /app/.env

# Optionally set environment variables directly (or override at runtime)
# ENV EXAMPLE_VAR=value

# Entry point for the container
CMD ["./app"]
