OAPI_CODEGEN_VERSION=v2.8.0
REDOC_CLI_VERSION=2.49.0

up:
# ビルドして起動
	make build && docker compose up -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

ps:
	docker compose ps

lint:
	make lint-web
	make lint-api

lint-web:
	cd web && pnpm run lint

lint-api:
	cd api && go vet $$(go list ./... | grep -v /pkg/interface/gen/)

format:
	make format-web
	make format-api

format-web:
	cd web && pnpm run format

format-api:
	cd api && go fmt $$(go list ./... | grep -v /pkg/interface/gen/)

test:
	make test-web && make test-api

test-web:
	cd web && pnpm run test

test-api:
	cd api && \
	set -a && source ../.env && \
	set +a && \
	go test ./... -count=1

test-e2e:
	cd web && \
	set -a && source ../.env && \
	set +a && \
	pnpm run test:e2e

gen:
	make gen-web
	make gen-api

gen-web:
	cd web && pnpm run gen

gen-api:
	npx --yes @redocly/cli@$(REDOC_CLI_VERSION) bundle openapi/openapi.yaml -o api/pkg/interface/gen/openapi/openapi.yaml
	cd api && \
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION) \
	-config ../openapi/oapi-codegen.yaml pkg/interface/gen/openapi/openapi.yaml

migrate:
	docker compose run --rm migrate

migrate-reset:
	docker compose run --rm -e FLYWAY_CLEAN_DISABLED=false migrate clean && make migrate
