FROM golang:1.18-alpine

RUN env GOOS=linux GOARCH=arm64 go build -o ./schoolplus ./cmd/http/main.go

COPY schoolplus /usr/local/bin/

CMD ["/usr/local/bin/schoolplus"]
