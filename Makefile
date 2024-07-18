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
app:
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