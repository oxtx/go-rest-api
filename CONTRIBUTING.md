# Contributing

## Prerequisites
- Go 1.22+
- Docker & docker compose
- golangci-lint
- migrate
- swag

## Setup
cp .env.example .env
docker compose up -d db
make migrate-up
make run

## Branching
feature/<short-desc>, fix/<short-desc>

## Commit Style
Conventional Commits (feat:, fix:, chore:, docs:, refactor:, test:)

## PR Checklist
- [ ] Tests added/updated
- [ ] Lint passes
- [ ] Swagger docs updated (make swag)
- [ ] No sensitive info
