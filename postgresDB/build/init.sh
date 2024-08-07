#!/bin/bash
set -e

# Run all SQL scripts in the /docker-entrypoint-initdb.d directory
for f in /docker-entrypoint-initdb.d/*.sql; do
    echo "Executing $f"
    psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$f"
done
