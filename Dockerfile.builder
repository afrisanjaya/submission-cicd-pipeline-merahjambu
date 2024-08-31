# Build Stage
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir -p /root/.cache/go-build && chown -R nobody:nogroup /root/.cache

USER nobody

COPY . .

RUN go test ./...

RUN go build -o main .
