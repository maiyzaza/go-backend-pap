# # Create build stage based on buster image
# FROM golang:1.19 AS builder
# # Create working directory under /app
# WORKDIR /app
# # Copy over all go config (go.mod, go.sum etc.)
# COPY go.* ./
# # Install any required modules
# RUN go mod download
# # Copy over Go source code
# COPY *.go ./
# # Run the Go build and output binary under hello_go_http
# RUN go build -o /hello_go_http
# # Make sure to expose the port the HTTP server is using
# EXPOSE 8080
# # Run the app binary when we run the container
# ENTRYPOINT ["/hello_go_http"]

# Create build stage based on buster image
FROM golang:latest

# Set the proper working directory within the Docker build context
WORKDIR /app

# Copy the entire project directory to the Docker build context
COPY . .

# Build the Go binary with vendored dependencies
RUN go build -o /hello_go_http

# Make sure to expose the port the HTTP server is using
EXPOSE 8080

# Run the app binary when we run the container
ENTRYPOINT ["/hello_go_http"]


# syntax=docker/dockerfile:1

# FROM golang:1.19

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8080

# # Run
# CMD ["/docker-gs-ping"]
