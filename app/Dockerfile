FROM golang:1.15.2-buster

EXPOSE 8080

WORKDIR /app

COPY http_server.go .

RUN go build -o http_server

ENTRYPOINT ["/app/http_server"]
