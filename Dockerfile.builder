# Build Stage
FROM golang:1.22.4 AS builder

WORKDIR /app

RUN mkdir -p /app/.cache/go-build && chmod -R 777 /app/.cache
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o main .
