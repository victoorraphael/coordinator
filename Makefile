export PORT = 8080
export MONGO_URI = mongodb://localhost:27017

run: compose-up
	go run cmd/http/main.go

local: compose-db
	go run cmd/http/main.go

compose-up: compose-down
	docker-compose up --build

compose-down:
	docker-compose down

compose-db:
	docker-compose up -d --build database