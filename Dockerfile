# Build Stage
FROM golang:1.19 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o cowsay-pro

# Runtime Stage
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/cowsay-pro .

CMD ["./cowsay-pro"]

