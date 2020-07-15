.PHONY: run
run:
	go run main.go web

.PHONY: start
start:
	docker-compose -f dockers/docker-compose.yml build && docker-compose -f dockers/docker-compose.yml up

.PHONY: lint
lint:
	golint ./...

.PHONY: gen_migration
gen_migration:
	go run migration/main.go new ${FILENAME}

.PHONY: up_migrate
up_migrate:
	go run migration/main.go up ${ARG}

.PHONY: down_migrate
down_migrate:
	go run migration/main.go down ${ARG}

.PHONY: drop_migrate
drop_migrate:
	go run migration/main.go drop

.PHONY: migrate_status
migrate_status:
	go run migration/main.go version

.PHONY: refresh_db
refresh_db:
	go run migration/main.go drop && \
	go run migration/main.go up

.PHONY: up_migrate_prod
up_migrate_prod:
	API_REVISION=release go run migration/main.go up
