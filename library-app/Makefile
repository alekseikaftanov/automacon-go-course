# Variables
MIGRATIONS_DIR = ./migrations
DB_URL = postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable
APP_NAME = library-app
GOOSE = goose

# Default target
all: build

# Migration targets
migrate-up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

migrate-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -dir $(MIGRATIONS_DIR) create $$name sql

# Build targets
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) ./cmd/...

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Test targets
test:
	@echo "Running tests..."
	go test ./...

# Run the application
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Show help
help:
	@echo "Makefile commands:"
	@echo "  make migrate-up       - Run database migrations up"
	@echo "  make migrate-down     - Run database migrations down"
	@echo "  make migrate-status   - Show migration status"
	@echo "  make migrate-create   - Create a new migration"
	@echo "  make build            - Build the application"
	@echo "  make clean            - Clean up build artifacts"
	@echo "  make test             - Run tests"
	@echo "  make run              - Run the application"

.PHONY: all migrate-up migrate-down migrate-status migrate-create build clean test run help