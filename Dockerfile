# Stage 1: Install dependencies
FROM golang:1.23.4-bookworm AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Stage 2: Build the application
FROM golang:1.23.4-bookworm AS builder
WORKDIR /app
COPY --from=deps /go/pkg /go/pkg
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOAMD64=v3

RUN go build -ldflags="-w -s" -o cmd/calculator/main ./cmd/calculator
RUN go build -o cmd/migrator ./cmd/migrator

# Stage 3: Run migrations and then the application
FROM debian:bullseye-slim  
WORKDIR /app
COPY --from=builder /app/cmd/calculator/main .
COPY --from=builder /app/cmd/migrator .
COPY local.env .
COPY config/local.yaml ./config/
COPY migrations ./migrations

RUN apt-get update && apt-get install -y bash

CMD ["sh", "-c", "./migrator --migrations-path=./migrations/postgres && ./main"]
