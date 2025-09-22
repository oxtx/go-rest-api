APP_NAME=go-rest-api
BIN=bin/$(APP_NAME)
PKG=./...
SWAG=github.com/swaggo/swag/cmd/swag@v1.8.12
LINT=github.com/golangci/golangci-lint/cmd/golangci-lint@latest

export GO111MODULE=on

.PHONY: run lint lint-fix tidy swag migrate-up-docker migrate-down-docker docker-up docker-down

run:
	go run ./cmd/server

lint:
	go run $(LINT) run

lint-fix:
	go run $(LINT) run --fix

tidy:
	go mod tidy

swag:
	go run $(SWAG) init -g cmd/server/main.go -o api/docs

migrate-up-docker:
	docker compose up migrate-up

migrate-down-docker:
	docker compose up migrate-down

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

clean:
	rm -rf api/docs
