ps:
	docker-compose ps
build:
	docker-compose up -d --build
up:
	docker-compose up -d
down:
	docker-compose down
restart:
	docker-compose restart
stop:
	docker-compose stop
go:
	docker-compose exec app sh
db:
	docker-compose exec db sh
installCLI:
	docker-compose exec app ./go_install.sh
install:
	docker-compose exec app go mod tidy
vendor:
	docker-compose exec app go mod vendor
dev:
	docker-compose exec app bee run

include .env
DATABASE_URL := "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE}&timezone=${DB_TIMEZONE}"
MIGRATIONS_PATH := database/migrations

migrate:
	docker-compose exec app bee migrate -driver=postgres -conn=$(DATABASE_URL) -dir="$(MIGRATIONS_PATH)"
migrate-new:
	docker-compose exec app bee generate migration $(name) -driver postgres
migrate-reset:
	docker-compose exec app bee migrate reset -driver=postgres -conn=$(DATABASE_URL) -dir="$(MIGRATIONS_PATH)"
migrate-rollback:
	docker-compose exec app bee migrate rollback -driver=postgres -conn=$(DATABASE_URL) -dir="$(MIGRATIONS_PATH)"
migrate-refresh:
	docker-compose exec app bee migrate refresh -driver=postgres -conn=$(DATABASE_URL) -dir="$(MIGRATIONS_PATH)"
