FROM golang:1.18-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=8080 \
#    DB_URI="postgres://jlsfrbek:41YTk-sbD-OIQyz7Odh9E7BIksrvHhu0@babar.db.elephantsql.com/jlsfrbek"
     DB_URI="postgres://root:secret@localhost:5454/schoolplus?sslmode=disable"

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build -o /app/main /app/cmd/main.go" --command="/app/main"