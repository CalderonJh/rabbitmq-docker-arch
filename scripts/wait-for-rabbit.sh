#!/bin/sh
set -e

host="$1"
shift
until nc -z "$host" 5672; do
  echo "Waiting for RabbitMQ at $host:5672..."
  sleep 3
done

exec "$@"