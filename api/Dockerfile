# Specifies a parent image
FROM golang:1.19.2-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copy the go module files and download dependenxies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
 
# Builds your app with optional configuration
RUN go build -o app main.go 
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD ["/bin/bash", "-c", "./app"]
#“./app” ]
