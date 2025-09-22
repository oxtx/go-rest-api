APP_NAME=go-rest-api
BIN=bin/$(APP_NAME)
PKG=./...
SWAG=github.com/swaggo/swag/cmd/swag@v1.8.12
LINT=github.com/golangci/golangci-lint/cmd/golangci-lint@latest

export GO111MODULE=on

.PHONY: run lint lint-fix tidy swag migrate-up-docker migrate-down-docker docker-up docker-down clean-swag clean-coverage

run:
	go run ./cmd/server

lint:
	go run $(LINT) run

lint-fix:
	go run $(LINT) run --fix
	
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

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

clean-swag:
	rm -rf api/docs

clean-coverage:
	rm -f coverage.out coverage.html

