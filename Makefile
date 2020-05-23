include .env
export

.PHONY: db_create db_start db_stop db_clean db_cli migrations migrations_down seeds run test

db_create:
	docker create --name db-api \
		-v dernacktehalloumi_db-api_data:/var/lib/postgresql/data \
		--env POSTGRES_PASSWORD=${DATABASE_PASSWORD} \
		--env POSTGRES_USER=${DATABASE_USER} \
		--env POSTGRES_DB=${DATABASE_NAME} \
		-p ${DATABASE_HOST}:${DATABASE_PORT}:5432/tcp \
		postgres:12

db_start:
	docker start db-api

db_stop:
	docker stop db-api

db_clean:
	docker rm -v db-api && docker volume rm dernacktehalloumi_db-api_data

db_cli:
	pgcli postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}

migrations:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.8.0 \
		-path=/migrations/ \
		-database postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable up

migrations_down:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.8.0 \
		-path=/migrations/ \
		-database postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable down -all

seeds:
	psql postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable -a -f seeds/initial.sql -f seeds/psql-tsv-cmds

run:
	go run cmd/rest/main.go

test:
	go test -v ./...
