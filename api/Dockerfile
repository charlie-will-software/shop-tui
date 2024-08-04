# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder

# Set COPY destination
WORKDIR /app

# Download GO Modules
COPY go.mod go.sum ./
RUN go mod download

# Step 5: Copy the source code into the container
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# OPTIONAL: Expose TCP port
# EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]
