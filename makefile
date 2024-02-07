.DEFAULT_GOAL := help

include .env
export

# === ENVIRONMENTS ============================================================
.PHONY: local-up
local-up: ## Start local environment
	docker-compose -f docker-compose.local.yml up --build -d

.PHONY: local-down
local-down: ## Down and clean up local environment
	@docker compose -f docker-compose.local.yml down -v

.PHONY: test-env-up
test-env-up: ## Run test environment.
	@docker compose -f docker-compose.test.yml up --exit-code-from migrate migrate

.PHONY: test-env-down
test-env-down: ## Down and cleanup test environment.
	@docker compose -f docker-compose.test.yml down -v

# === TESTS ===================================================================
.PHONY: test
test: ## Run all tests.
	@go test -race -count=1 -coverprofile=cover.out ./...
	@go tool cover -func=cover.out | grep ^total | tr -s '\t'

.PHONY: test-short
test-short: ## Run only unit tests, tests without I/O dependencies.
	@go test -short -coverprofile=cover.out ./...

.PHONY: cover
cover: ## Show coverage in browser
	@go tool cover -html=cover.out

# === MIGRATIONS ==============================================================
.PHONY: create-migration
create-migration: ## Create migration named $name
	@migrate create -ext sql -dir migrations/ -seq -digits 3 $(name)

.PHONY: migrate-up
migrate-up:	## Migrate up $(V) local postgres DB
	@migrate -source file://./migrations -database "postgres://localhost:5432/appolodb?user=appolo&password=$(POSTGRES_PASSWORD)&sslmode=disable" up $(V)

.PHONY: migrate-down
migrate-down:	## Migrate down $(V) local postgres DB 
	@migrate -source file://./migrations -database "postgres://localhost:5432/appolodb?user=appolo&password=$(POSTGRES_PASSWORD)&sslmode=disable" down $(V)

.PHONY: migrate-force
migrate-force:	## Force migration version to $(V)
	@migrate -source file://./migrations -database "postgres://localhost:5432/appolodb?user=appolo&password=$(POSTGRES_PASSWORD)&sslmode=disable" force $(V)

# === GENERAL =================================================================
.PHONY: run
run: ## Start the service with the CONFIG_PATH configuration
	@go run cmd/appolo/appolo.go

.PHONY: help
help: ## Show help for each of the Makefile targets.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
