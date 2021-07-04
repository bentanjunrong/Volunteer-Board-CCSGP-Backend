#!/usr/bin/env bash

# Show env vars
grep -v '^#' .env

# Export env vars
export $(grep -v '^#' .env | xargs)

# Update vals in monstache.toml
sed -i s/DB_URL/$DB_URL/g monstache.toml
sed -i s/ES_PASSWORD/$ES_PASSWORD/g monstache.toml

{
    echo "[Service]"; 
    grep -v '^#' .env | while read -r line ; do
        echo "Environment=\"$line\""
    done
} | sudo tee /etc/systemd/system/volunteery.service.d/override.conf

{
    echo "[Service]"; 
    grep -v '^#' .env | while read -r line ; do
        echo "Environment=\"$line\""
    done
} | sudo tee /etc/systemd/system/monstache.service.d/override.conf
