package generator

import (
	"os"
	"path/filepath"
)

func writeMakefile(cfg Config) error {
	projectRoot := cfg.ProjectName

	makefileContent := `# Simple Makefile for GoCozyy project

PROJECT_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# Build the application
build:
	@echo "Building..."
	@cd $(PROJECT_DIR)/backend && go build -o main main.go

# Run the application
run:
	@cd $(PROJECT_DIR)/backend && go run main.go

# Create DB container
docker-run:
	@cd $(PROJECT_DIR) && \
	if [ -f docker-compose.db.yml ]; then \
		echo "Starting PostgreSQL..."; \
		docker compose -f docker-compose.db.yml up -d; \
	elif [ -f docker-compose.yml ]; then \
		echo "Starting full stack..."; \
		docker compose up -d; \
	else \
		echo "No docker-compose file found. Run with --docker flag or use postgres DB."; \
		exit 1; \
	fi

# Shutdown DB container
docker-down:
	@cd $(PROJECT_DIR) && \
	if [ -f docker-compose.db.yml ]; then \
		docker compose -f docker-compose.db.yml down; \
	elif [ -f docker-compose.yml ]; then \
		docker compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@cd $(PROJECT_DIR)/backend && go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f $(PROJECT_DIR)/backend/main

# Live Reload (Go)
watch:
	@cd $(PROJECT_DIR)/backend && \
	if command -v air > /dev/null 2>&1; then \
		air; \
		echo "Watching...";\
	else \
		read -p "Go's 'air' is not installed. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
		else \
			echo "Skipping air install."; \
			exit 1; \
		fi; \
	fi

.PHONY: all build run test clean watch docker-run docker-down
`
	return os.WriteFile(filepath.Join(projectRoot, "Makefile"), []byte(makefileContent), 0o644)
}
