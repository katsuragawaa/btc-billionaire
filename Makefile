.PHONY: migrate_force migrate_version migrate_up migrate_down docker local swaggo test

# ==============================================================================
# Migrate postgresql

migrate_force:
	migrate -database postgres://postgres:postgres@localhost:5432/btc_db?sslmode=disable -path migrations force 1

migrate_version:
	migrate -database postgres://postgres:postgres@localhost:5432/btc_db?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/btc_db?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/btc_db?sslmode=disable -path migrations down 1


# ==============================================================================
# Docker compose commands

docker:
	echo "Starting docker environment"
	docker-compose -f docker-compose.dev.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build


# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go


# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

client:
	go run ./cmd/client/main.go


# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)