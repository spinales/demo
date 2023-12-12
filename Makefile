SHELL=/bin/bash

run:
	go run ./cmd/demo/

build:
	templ generate
	go build -o ./tmp/demo ./cmd/demo
	./tmp/demo

compose:
	templ generate
	pack build demo --env-file .env --builder paketobuildpacks/builder-jammy-tiny
	docker compose -f docker-compose.yml up

dev:
	docker compose -f docker-compose-dev.yml up && ./tmp/demo