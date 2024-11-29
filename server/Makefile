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
