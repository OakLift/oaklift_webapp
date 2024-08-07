#!/bin/bash

set -euo pipefail

POSTGRES_DB_DIR="../postgresDB"

setupPostgresDB() {
    mkdir -p "$POSTGRES_DB_DIR/output"
    cp "$POSTGRES_DB_DIR/build/postgres.yml" "$POSTGRES_DB_DIR/output/postgres.yml"
    cp "$POSTGRES_DB_DIR/build/users.sql" "$POSTGRES_DB_DIR/output/users.sql"
    cp "$POSTGRES_DB_DIR/build/default.env" "$POSTGRES_DB_DIR/output/.env"
}

startPostgresDB() {
    docker compose -f "$POSTGRES_DB_DIR/output/postgres.yml" up -d oak-db
}

initPostgresDB() {
    setupPostgresDB
    startPostgresDB
}

start() {
    initPostgresDB
}

if [[ "$1" == "start" ]]; then
    start
fi

if [[ "$1" == "stop" ]]; then
    docker kill "$(docker ps -aq)" || true
    docker rm "$(docker ps -a -q)" || true
    docker network prune -f
    docker volume prune -f
fi



