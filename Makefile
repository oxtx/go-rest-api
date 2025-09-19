APP_NAME=go-rest-api
BIN=bin/$(APP_NAME)
PKG=./...
SWAG=github.com/swaggo/swag/cmd/swag

export GO111MODULE=on

.PHONY: all build run test lint tidy swag migrate-up migrate-down docker-up docker-down clean coverage

all: build

build:
	mkdir -p bin
	go build -o $(BIN) ./cmd/server

run:
	go run ./cmd/server

test:
	go test -race -coverprofile=coverage.out $(PKG)

coverage: test
	go tool cover -func=coverage.out | grep total

lint:
	golangci-lint run

tidy:
	go mod tidy

swag:
	swag init -g cmd/server/main.go -o api/docs

migrate-up:
	migrate -path internal/db/migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path internal/db/migrations -database "$$DATABASE_URL" down 1

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

clean:
	rm -rf bin coverage.out
