#!/bin/sh

set -e
echo "Waiting for database to be ready at $MYSQL_HOST:$MYSQL_PORT..."

echo "Database is ready."

# Run the migrations
echo "Running database migrations..."
migrate -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(restro_db:3306)/${MYSQL_DATABASE}" -path "./database/migrations" up

echo "Migrations applied successfully."

# This will run the CMD from your Dockerfile (e.g., /main)
exec "$@"