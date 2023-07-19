FROM golang:1.19

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["/app/main"]


# # Create build stage based on buster image
# FROM golang:latest

# # Set the proper working directory within the Docker build context
# WORKDIR /app

# # Copy the entire project directory to the Docker build context
# COPY . .

# # Build the Go binary with vendored dependencies
# RUN go build -o /hello_go_http

# # Make sure to expose the port the HTTP server is using
# EXPOSE 8080

# # Run the app binary when we run the container
# ENTRYPOINT ["/hello_go_http"]
