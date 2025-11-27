.PHONY: create-migration

MIGRATIONS_DIR=./db/migrations

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

generate-models:
	oapi-codegen -package models -generate types -o internal/models/models.gen.go specs/openapi.yaml