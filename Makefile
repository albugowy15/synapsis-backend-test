DB_URL=postgres://postgres:postgres@localhost:5432/synapsis_db?sslmode=disable

network:
	docker network create synapsis-network

postgres:
	docker run --name postgres --network synapsis-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres synapsis_db

dropdb:
	docker exec -it postgres dropdb synapsis_db

migrate_up:
	migrate -path database/migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path database/migrations -database "$(DB_URL)" -verbose down

migrate_fix:
	migrate -path database/migrations -database postgres://postgres:postgres@localhost:5432/synapsis_db force ${VERSION}

new_migration:
	migrate create -ext sql -dir database/migrations -seq $(name)

seed:
	go run cmd/seeder.go

test:
	go test -v -cover -short ./...

build_api:
	docker build -t synapsis-backend-test .

server_api:
	docker run --name synapsis-backend-test --network synapsis-network -p 8080:8080 -e DB_SOURCE="postgres://postgres:postgres@postgres:5432/synapsis_db?sslmode=disable" synapsis-backend-test 

server:
	go run main.go

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

build:
	go build -o tmp/main main.go

.PHONY: network postgres createdb dropdb migrate_up migrate_down migrate_fix nem_migration seed test build_api server_api server redis build
