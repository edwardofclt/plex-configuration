#!/bin/bash

# Stop stacks in reverse order, only one level deep
find /plex -mindepth 2 -maxdepth 2 -type f \( -name "docker-compose.yml" -o -name "docker-compose.yaml" \) | tac | while read -r compose_file; do
    dir=$(dirname "$compose_file")
    echo "Stopping Docker Compose in $dir"
    (cd "$dir" && docker compose down)
done

docker compose down
