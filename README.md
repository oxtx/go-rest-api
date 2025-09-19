# go-rest-api

Minimal, production-ready REST API scaffold in Go (Gin + PostgreSQL) with Swagger, migrations, logging, and basic structure.

## Features
- Gin HTTP server
- Layered architecture (handler → service → repository)
- PostgreSQL + migrations (golang-migrate)
- Swagger (swag) docs at /swagger/index.html
- Request ID + basic logging (slog)
- Makefile automation
- Docker & docker compose

## Tech
Go 1.25, Gin, pgx, swag, testify, slog, golangci-lint.

## Structure (essentials)
cmd/server/main.go         # entrypoint
internal/
handler/                 # HTTP handlers
service/                 # business logic
repository/              # DB queries
dto/                     # request/response structs
model/                   # domain models
server/                  # router + middleware
db/migrations/           # SQL migrations
pkg/response/              # response helpers
scripts/gen_openapi.sh     # swagger generation

## Quick Start
```bash
git clone <repo-url> go-rest-api
cd go-rest-api
cp .env.example .env
docker compose up -d db
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/awesome?sslmode=disable"
make migrate-up
make swag    # generate api/docs
make run
# Health check
curl http://localhost:8080/healthz