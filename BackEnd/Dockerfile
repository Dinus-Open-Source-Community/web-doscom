# stage 1: build
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o  app ./cmd/api/main.go

# stage 2: run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

ENV GO_ENV=development
ENV PORT=8080
