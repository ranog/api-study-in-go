include .env

create-migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq init

migrate-up:
	migrate -path=internal/database/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up

migrate-down:
	migrate -path=internal/database/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" down

.PHONY: create-migration migrate-up migrate-down
