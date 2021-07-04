#!/usr/bin/env bash

# Show env vars
grep -v '^#' .env

# Export env vars
export $(grep -v '^#' .env | xargs)

# Update vals in monstache.toml
sed -i s/DB_URL/$DB_URL/g monstache.toml
sed -i s/ES_PASSWORD/$ES_PASSWORD/g monstache.toml
