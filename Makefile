# Define variables for convenience
DOCKER_COMPOSE=docker-compose
DOCKER=docker

# Define service names
APP_SERVICE=app
DB_SERVICE=db

# Define Docker image names
APP_IMAGE=your-app-image
DB_IMAGE=postgres:13

# Define Docker Compose file
COMPOSE_FILE=docker-compose.yml

# Define Docker Compose network
NETWORK=app-network

# Build Docker images
build:
	@echo "Building Docker images..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

# Start services with Docker Compose
up:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d

# Stop services with Docker Compose
down:
	@echo "Stopping services..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

# Remove Docker containers
clean:
	@echo "Removing Docker containers..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down -v

# Remove Docker images
clean-images:
	@echo "Removing Docker images..."
	$(DOCKER) rmi $(APP_IMAGE) $(DB_IMAGE)

# View logs
logs:
	@echo "Viewing logs..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f

# Run bash in the app container
bash:
	@echo "Running bash in the app container..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) exec $(APP_SERVICE) /bin/sh

# Initialize database (useful for running SQL scripts manually)
init-db:
	@echo "Initializing database..."
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) exec $(DB_SERVICE) psql -U postgres -d yourdbname -f /docker-entrypoint-initdb.d/users.sql

# Print help message
help:
	@echo "Available commands:"
	@echo "  build         - Build Docker images"
	@echo "  up            - Start services with Docker Compose"
	@echo "  down          - Stop services with Docker Compose"
	@echo "  clean         - Remove Docker containers"
	@echo "  clean-images  - Remove Docker images"
	@echo "  logs          - View logs"
	@echo "  bash          - Run bash in the app container"
	@echo "  init-db       - Initialize database with SQL script"
	@echo "  help          - Print this help message"
