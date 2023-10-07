include .env

# Variables
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BUILD_FILE = release/go-next-docker
MAIN_FILE = cmd/main.go

# Project commands
build:
	$(GOBUILD) -o $(BUILD_FILE) $(MAIN_FILE)

test:
	$(GOTEST) -v ./...

clean:
	rm -f $(BUILD_FILE)

run:
	air

lint:
	@echo "🏃..."
	golangci-lint run

.PHONY: build test clean run lint

# Database commands
db_create:
	docker exec -it postgres psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);"

db_drop:
	docker exec -it postgres psql -U $(DB_USER) -c "DROP DATABASE $(DB_NAME);"

.PHONY: db_create db_drop

# Model commands
MODEL_DIR = internal/app/models

generate:
	@if [ -n "$(model)" ]; then \
		touch $(MODEL_DIR)/$(model).go; \
		echo "🚀... \"$(model).go\" generated in \"$(MODEL_DIR)\""; \
		echo "package models\n\ntype $(model) struct {}" > $(MODEL_DIR)/$(model).go; \
		echo "Don't forget to add a model into \"migrations.go\" 👍"; \
	else \
		echo "'name' is not defined. Please provide a name to generate a file, e.g., make generate_model name=mymodel"; \
	fi

.PHONY: generate_model

# Do seeds data
SEED_DIR = internal/database/seed
seed:
	go run $(SEED_DIR)/$(model).go

# Help for displaying available commands
help:
	@echo "💡 Available commands:"
	@echo " - make build"
	@echo " - make test"
	@echo " - make clean"
	@echo " - make run"
	@echo " - make lint"
	@echo " - make db_create"
	@echo " - make db_drop"
	@echo " - make generate model=user"
	@echo " - make seed model=post"
