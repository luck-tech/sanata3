BINARY_NAME=app
ENV_FILE=

setup:
	cp .env.example .env

build:
	docker compose build

run: build
	docker compose up

rund: build
	docker compose up -d

test:
	go test -v ./...

name=
new-migrate:
	migrate create -ext sql -dir cmd/migrate/source -seq $(name)

migrate:
	go run cmd/migrate/main.go -e .env.migrate -f cmd/migrate/source

swag:
	swag init --dir cmd/app/ --parseDependencyLevel 10 