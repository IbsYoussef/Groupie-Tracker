#!/bin/bash
# scripts/run-migrations.sh
# Runs all migration files in scripts/migrations/ in numerical order

set -e

MIGRATIONS_DIR="/docker-entrypoint-initdb.d/migrations"
PGUSER="${POSTGRES_USER:-groupie_user}"
PGDB="${POSTGRES_DB:-groupie_tracker}"

echo "🔄 Running database migrations..."

# Sort and run all .sql files
for migration in $(ls -1 $MIGRATIONS_DIR/*.sql 2>/dev/null | sort -V); do
    filename=$(basename "$migration")
    echo "  ⚙️  Applying: $filename"
    psql -U "$PGUSER" -d "$PGDB" -f "$migration"
done

echo "✅ All migrations applied successfully"