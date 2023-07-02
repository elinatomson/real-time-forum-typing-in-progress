#!/bin/bash

# Clear terminal screen
clear

# Building docker image
docker image build -t realtimeforum .

# Run a command in a new container and get the container ID
CONTAINER_ID=$(docker container run -p 8080:8080 -d realtimeforum)

# Display the container ID
echo "Container ID: $CONTAINER_ID"

# Prompt user for cleanup
read -p "Open http://localhost:8080 to visit the forum. When you are finished, then press Enter to run delete_docker script"

# Run the cleanup script passing the container ID
./delete_docker.sh "$CONTAINER_ID"