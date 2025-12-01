.PHONY: create-migration generate-models db-up db-down migrate-up migrate-down start stop docker-build

MIGRATIONS_DIR=./db/migrations
DB_DSN="postgresql://dev:dev@localhost:5432/dev"

#------------MIGRATIONS------------

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

#------------DOCKER------------

docker-compose-build:
	docker-compose build --no-cache

docker-compose-up:
	docker-compose up -d
	@sleep 2

docker-compose-down:
	docker-compose down -v
	@sleep 2

docker-build:
	docker build -t api-question-service -f Dockerfile .

#------------SERVICE------------

start: docker-compose-build docker-compose-up migrate-up
	@echo "Сервис запущен!"

stop: docker-compose-down
	@echo "Сервис остановлен!"

restart: docker-compose-down docker-compose-up migrate-up
	@echo "Сервис перезапущен!"