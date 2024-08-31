# Build Stage
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main /main

EXPOSE 8181

# Run the application
ENTRYPOINT ["/main"]
