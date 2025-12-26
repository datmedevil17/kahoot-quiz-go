.PHONY: run-server run-client build-server build-client docker-up docker-down install

# Server
run-server:
	cd server && go run ./cmd/api

build-server:
	cd server && go build -o bin/api ./cmd/api

# Client
run-client:
	cd client && npm run dev

build-client:
	cd client && npm run build

install:
	cd server && go mod download
	cd client && npm install

# Docker
docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f
