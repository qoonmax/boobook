include .env

to-dev:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest
	docker compose up -d --build
	$(MAKE) migrate_up

run:
	docker compose up -d --build
	go run cmd/api/main.go

air:
	docker compose up -d --build
	air server --port 8000

create_migration:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose create $(name) sql

migrate_up:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose up

migrate_down:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose down

migrate_fresh:
	docker compose exec -T postgres psql -U $(DATABASE_USER) -d $(DATABASE_NAME) -c 'DROP SCHEMA public CASCADE; CREATE SCHEMA public;'