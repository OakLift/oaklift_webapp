#!/bin/bash

# Stop all running containers
echo "Stopping all running containers..."
running_containers=$(docker ps -q)
if [ -n "$running_containers" ]; then
    docker stop $running_containers
else
    echo "No running containers found."
fi

# Remove all containers
echo "Removing all containers..."
all_containers=$(docker ps -aq)
if [ -n "$all_containers" ]; then
    docker rm $all_containers -f
else
    echo "No containers found."
fi

# Remove all volumes
echo "Removing all volumes..."
all_volumes=$(docker volume ls -q)
if [ -n "$all_volumes" ]; then
    docker volume rm $all_volumes
else
    echo "No volumes found."
fi

echo "All containers and volumes have been removed."