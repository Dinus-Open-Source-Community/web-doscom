# stage 1: build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/api/main.go
RUN CGO_ENABLED=0 go build -o migrate ./cmd/migrate/main.go
RUN CGO_ENABLED=0 go build -o seeder ./cmd/seeder/main.go

# stage 2: run
FROM alpine:3.20.0

WORKDIR /app

# Install postgresql-client for DB wait check
RUN apk add --no-cache postgresql-client

COPY --from=builder /app/app .
# COPY --from=builder /app/storage/uploads/foto ./storage/uploads/foto
COPY --from=builder /app/migrate .
COPY --from=builder /app/seeder .
COPY --from=builder /app/migrations ./migrations
COPY docker-entrypoint.sh .
COPY .env .
RUN chmod +x docker-entrypoint.sh

ENV GO_ENV=development
ENV PORT=8080
ENTRYPOINT ["./docker-entrypoint.sh"]
