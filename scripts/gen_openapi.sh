#!/usr/bin/env bash
set -euo pipefail
swag init -g cmd/server/main.go -o api/docs
