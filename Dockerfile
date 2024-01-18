FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todos ./cmd

# Path: Dockerfile
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/todos .

EXPOSE 8080

CMD ["./todos"]
