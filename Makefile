APP_NAME=go-rest-api
BIN=bin/$(APP_NAME)
PKG=./...
SWAG=github.com/swaggo/swag/cmd/swag@v1.8.12
LINT=github.com/golangci/golangci-lint/cmd/golangci-lint@latest

export GO111MODULE=on

.PHONY: run lint lint-fix tidy swag migrate-up-docker migrate-down-docker docker-up docker-down clean-swag clean-coverage test test-coverage fmt fmt-check ci-local

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

# New CI-related commands
fmt:
	gofumpt -w .

fmt-check:
	@if [ "$$(gofumpt -l . | wc -l)" -gt 0 ]; then \
		echo "The following files are not formatted:"; \
		gofumpt -l .; \
		echo "Run 'make fmt' to fix formatting"; \
		exit 1; \
	fi

tidy-check:
	@go mod tidy
	@if [ -n "$$(git status --porcelain go.mod go.sum)" ]; then \
		echo "go.mod or go.sum is not tidy"; \
		git diff go.mod go.sum; \
		exit 1; \
	fi

# Run all CI checks locally
ci-local: swag fmt-check tidy-check lint test-coverage
	@echo "✅ All CI checks passed!"

# Install development tools
install-tools:
	go install $(SWAG)
	go install $(LINT)
	go install mvdan.cc/gofumpt@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
