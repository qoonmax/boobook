include .env

to-dev:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	docker compose up -d

run:
	go run cmd/api/main.go

create_migration:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose create $(name) sql

migrate_up:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose up

migrate_down:
	GOOSE_DRIVER=$(DATABASE_DRIVER) GOOSE_DBSTRING=$(DATABASE_DBSTRING) GOOSE_MIGRATION_DIR=$(DATABASE_MIGRATION_DIR) goose down

migrate_fresh:
	docker compose exec -T postgres psql -U $(DATABASE_USER) -d $(DATABASE_NAME) -c 'DROP SCHEMA public CASCADE; CREATE SCHEMA public;'