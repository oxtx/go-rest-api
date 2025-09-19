#!/usr/bin/env bash
set -e

HOST=${1:-localhost}
PORT=${2:-5432}
TIMEOUT=${3:-30}

echo "Waiting for $HOST:$PORT up to $TIMEOUT seconds..."
for i in $(seq 1 "$TIMEOUT"); do
  if nc -z "$HOST" "$PORT" 2>/dev/null; then
    echo "Service available."
    exit 0
  fi
  sleep 1
done
echo "Timeout waiting for $HOST:$PORT"
exit 1
