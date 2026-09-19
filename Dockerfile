# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/cbog ./cmd/bot/main.go

# Final minimal runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates git

COPY --from=builder /app/bin/cbog /usr/local/bin/cbog

ENTRYPOINT ["/usr/local/bin/cbog"]
