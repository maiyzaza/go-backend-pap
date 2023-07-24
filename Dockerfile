FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 5000
CMD ["/app/main"]