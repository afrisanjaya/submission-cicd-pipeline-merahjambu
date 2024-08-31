# Build Stage
FROM golang:1.22.4 AS builder

WORKDIR /app

ENV GOCACHE=/app/.cache/go-build
ENV GOPATH=/app/go

RUN mkdir -p /app/.cache/go-build /app/go && chmod -R 777 /app/.cache

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o main .
