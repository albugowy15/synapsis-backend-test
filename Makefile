migrate-up: 
	migrate -path database/migrations -database postgresql://postgres:postgres@localhost:5432/synapsis_db up

migrate-down: 
	migrate -path database/migrations -database postgresql://postgres:postgres@localhost:5432/synapsis_db down

migrate-fix: 
	migrate -path database/migrations -database postgresql://postgres:postgres@localhost:5432/synapsis_db force ${VERSION}

seed:
	go run cmd/seeder.go

build:
	go build -o tmp/main main.go

.PHONY: migrate-up migrate-down migrate-fix seed build
