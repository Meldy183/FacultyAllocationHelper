FROM golang:1.24.3-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /service ./cmd/profile-service/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /service .
COPY --from=builder /app/config ./config

EXPOSE 8083
CMD ["./service"]