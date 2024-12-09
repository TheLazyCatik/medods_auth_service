#!/bin/bash

BASE_DIR="backend/db/init"
SCHEMA_FILE="db/schema.sql"
OUTPUT_DIR="backend/infra/repository/private/db/build"
SQLC_CONFIG="backend/infra/repository/private/db/sqlc.yaml"

mkdir -p "$OUTPUT_DIR"
cp "$SCHEMA_FILE" "$OUTPUT_DIR"

sqlc generate -f "$SQLC_CONFIG"