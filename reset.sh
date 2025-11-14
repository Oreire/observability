#!/bin/bash

echo "🔧 Stopping all running containers..."
docker stop $(docker ps -q)

echo "🗑️ Removing all containers..."
docker rm $(docker ps -aq)

echo "🧼 Pruning unused images, volumes, and networks..."
docker system prune -a --volumes -f

echo "✅ Docker environment reset complete. All ports are now free."

