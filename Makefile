OAPI_CODEGEN_VERSION=v2.8.0
REDOC_CLI_VERSION=2.49.0

up:
# ビルドして起動
	docker compose build && docker compose up -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

ps:
	docker compose ps

lint:

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