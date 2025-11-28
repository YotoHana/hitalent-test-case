.PHONY: create-migration

MIGRATIONS_DIR=./db/migrations
DB_DSN="postgresql://dev:dev@localhost:5432/dev"

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

generate-models:
	oapi-codegen -package models -generate types -o internal/models/models.gen.go specs/openapi.yaml

db-up:
	docker-compose up -d
	@echo "Ожидание запуска PostgreSQL..."
	@sleep 3

db-down:
	docker-compose down
	@echo "Ожидание остановки PostgreSQL..."
	@sleep 3

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

start: db-up migrate-up
	@echo "База данных готова!"