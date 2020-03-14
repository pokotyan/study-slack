.PHONY: run
run:
	go run main.go

.PHONY: start
start:
	docker-compose -f dockers/docker-compose.dev.yml build && docker-compose -f dockers/docker-compose.dev.yml up

.PHONY: lint
lint:
	golint ./...

.PHONY: gen_migration
gen_migration:
	go run migration/main.go new ${name}

.PHONY: up_migrate
up_migrate:
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go up ${arg}

.PHONY: down_migrate
down_migrate:
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go down ${arg}

.PHONY: drop_migrate
drop_migrate:
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go drop

.PHONY: migrate_status
migrate_status:
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go version

.PHONY: refresh_db
refresh_db:
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go drop && \
	docker-compose -f dockers/docker-compose.dev.yml run app go run migration/main.go up