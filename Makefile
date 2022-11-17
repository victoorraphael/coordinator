export PORT = 8080
export DB_URI = postgres://root:secret@localhost:5432/schoolplus?sslmode=disable

run: compose-up
	go run cmd/http/main.go

local: compose-db
	go run cmd/http/main.go

compose-up: compose-down
	docker-compose up --build

clear:
	docker-compose down

compose-db:
	docker-compose up -d --build database