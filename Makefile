DB_NAME=project-template-db
DB_PORT=35232
DB_PASSWORD=s0m3C0mpl3xP4ss
DB_IMAGE=postgres:alpine
DB_VOLUME=$(shell pwd)/db-data

.PHONY: api api-install api-build api-update db db-logs db-stop db-remove gen-types fe fe-install fe-build fe-update prod-build prod-publish prod-locally prod-locally-logs prod-locally-stop install update build
%:
	@:

# ---------- Backend ----------
api:
	cd ./api && air

api-install:
	cd ./api && go mod tidy

api-build:
	cd ./api && go build -o ./tmp/main .

api-update:
	cd ./api && go get -u all && go mod tidy && gofmt -w -l .

# ---------- Database ----------
db:
	@if [ $$(docker ps -a -q -f name=$(DB_NAME)) ]; then \
		echo "Starting existing database container..."; \
		docker start $(DB_NAME); \
	else \
		echo "Running database container for the first time..."; \
		docker run --name $(DB_NAME) -p $(DB_PORT):5432 -e POSTGRES_PASSWORD=$(DB_PASSWORD) -v $(DB_VOLUME):/var/lib/postgresql/data -d $(DB_IMAGE) -c max_connections=200; \
	fi

db-logs:
	docker logs $(DB_NAME) --follow
	
db-stop:
	@echo "Stopping database container..."
	docker stop $(DB_NAME)

db-remove:
	@echo "Removing database container..."
	docker stop $(DB_NAME)
	docker rm -f $(DB_NAME)

gen-types:
	rm -rf ./shared/fe/api-client/src/generated
	openapi-generator generate \
		-i ./api/generated/swagger.yaml \
		-g typescript-fetch \
		-o ./shared/fe/api-client/src/generated \
		--skip-validate-spec \
		--additional-properties=supportsES6=true,typescriptThreePlus=true,npmVersion=$$(npm --version)
	cd ./shared/fe/api-client && \
	bun install && \
	bun run build

# ---------- Frontend ----------
fe:
	cd ./fe && bun run dev

fe-install:
	cd ./fe && bun install

fe-build:
	cd ./fe && bun run build

fe-update:
	cd ./fe && bun update --latest
	cd ./shared/fe/api-client && bun update --latest
	cd ./fe && bun run lint
	
# ---------- Docker build for production ----------
VERSION := $(word 2,$(MAKECMDGOALS))
DOCKER_TAG = v$(shell echo $(VERSION) | sed 's/^v*//')
prod-build:
ifndef VERSION
	$(error VERSION is required. Usage: make prod-build vX.X.X)
endif
	@echo "Generating TypeScript types..."
	$(MAKE) gen-types
	@echo "Building Docker image $(DOCKER_TAG)..."
	cp .gitignore .dockerignore
	docker build -f Dockerfile.app --no-cache -t docker-registry.dev.tomasdiblik.cz/project-template:$(DOCKER_TAG) .

LOCAL_API_PROD_PORT ?= 35230
prod-locally:
ifndef VERSION
	$(error VERSION is required. Usage: make prod-locally vX.X.X)
endif
	@echo "Ensuring database is running..."
	$(MAKE) db
	$(MAKE) prod-locally-stop
	@echo "Running Docker image $(DOCKER_TAG) locally on port $(LOCAL_API_PROD_PORT)..."
	@docker run -d --env-file api/.env.production -p $(LOCAL_API_PROD_PORT):35230 docker-registry.dev.tomasdiblik.cz/project-template:$(DOCKER_TAG)

prod-locally-logs:
	@CONTAINER_ID=$$(docker ps -q --filter "publish=$(LOCAL_API_PROD_PORT)"); \
	if [ -z "$$CONTAINER_ID" ]; then \
		echo "No running container found on port $(LOCAL_API_PROD_PORT)."; \
	else \
		docker logs -f $$CONTAINER_ID; \
	fi
	
prod-locally-stop:
	@echo "Stopping local prod container, if it exists..."
	@docker ps -q --filter "publish=$(LOCAL_API_PROD_PORT)" | xargs -r docker stop

# ---------- Combined Targets ----------
install: api-install fe-install gen-types
update: api-update fe-update install build
build: api-build fe-build