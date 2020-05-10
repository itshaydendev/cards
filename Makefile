GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_NAME=cards
MIGRATIONS_DIR=migrations

# Default Task
all: build

# Build Tasks
build:
	$(GOBUILD) -o build/$(BINARY_NAME) -v github.com/itshaydendev/cards/cmd/cards

run:
	$(GORUN) cmd/cards/main.go

# Migration Tasks
migrate: migrate_up

migrate_up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up
migrate_down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down
migrate_new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)