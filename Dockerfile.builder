# Build Stage
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /root/.cache/go-build && chmod -R 777 /root/.cache

RUN go test ./...

RUN go build -o main .
