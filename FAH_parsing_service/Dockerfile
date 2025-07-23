FROM golang:1.23.10-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /service ./cmd/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /service .
COPY --from=builder /app/config ./config

EXPOSE 8082
CMD ["./service"]