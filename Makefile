run: compose-up
	go run cmd/http/main.go

compose-up: compose-down
	docker-compose up --build

compose-down:
	docker-compose down