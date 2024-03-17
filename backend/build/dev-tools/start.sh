#!/bin/bash
set -euo pipefail

DOCKER_COMPOSE_NEXUS_DB="../../nexusdb/build/nexusdb.yml"
DOCKER_COMPOSE_USER_GUARD="../../user-guard/build/user-guard.yml"

init_user_services() {
    echo "Initializing user services..."
    start_user_guard
}

init_databases() {
    echo "Initializing databases..."
    start_nexus_db
}

start_user_guard() {
    docker compose -f "$DOCKER_COMPOSE_USER_GUARD" up -d
}

start_nexus_db() {
    docker compose -f "$DOCKER_COMPOSE_NEXUS_DB" up -d
}

init_dev() {
    init_databases
    init_user_services
}

init_dev