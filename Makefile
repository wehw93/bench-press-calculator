.PHONY: build
build:
	go build -v ./cmd/calculator

.PHONY: runMigrations
runMigrations:
	go run ./cmd/migrator --migrations-path=./migrations/postgres

.PHONY: test

test:
	go test -v -race -timeout 30s ./...
	
.DEFAULT_GOAL := build
