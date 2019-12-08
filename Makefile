run:
	go run main.go
start:
	docker-compose -f docker-compose.dev.yml build && docker-compose -f docker-compose.dev.yml up
lint:
	golint ./...