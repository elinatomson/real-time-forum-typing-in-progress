#!/bin/bash

# Get the container ID passed as an argument
CONTAINER_ID="$1"

# Stop the container
docker container stop "$CONTAINER_ID"

# Remove the specific container
docker container rm "$CONTAINER_ID"