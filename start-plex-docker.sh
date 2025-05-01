#!/bin/bash

# Find docker-compose files only one level deep under /plex
find /plex -mindepth 2 -maxdepth 2 -type f \( -name "docker-compose.yml" -o -name "docker-compose.yaml" \) | while read -r compose_file; do
    dir=$(dirname "$compose_file")
    echo "Starting Docker Compose in $dir"
    (cd "$dir" && docker compose pull && docker compose up -d)
done

echo "Starting Docker Compose in /plex"
docker compose pull
docker compose up -d
