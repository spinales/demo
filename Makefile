SHELL=/bin/bash

run:
	docker compose -f docker-compose.yml up || { docker compose -f docker-compose.yml down; exit 0; }

build:
	templ generate
	go build -o ./tmp/demo ./cmd/demo
	pack build demo --env-file .env --builder paketobuildpacks/builder-jammy-tiny

dev:
	templ generate
	go build -o ./tmp/demo ./cmd/demo
	pack build demo --env-file .env --builder paketobuildpacks/builder-jammy-tiny
	docker compose -f docker-compose.yml up || { docker compose -f docker-compose.yml down; exit 0; }

test:
	docker compose -f docker-compose-dev.yml up -d --remove-orphans
	go test -v -cover ./internal/storage/... || { docker compose -f docker-compose-dev.yml down; exit 1; }
	docker compose -f docker-compose-dev.yml down
