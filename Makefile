DB_NAME=project-template_db
DB_PORT=35232
DB_PASSWORD=s0m3C0mpl3xP4ss
DB_IMAGE=postgres:alpine
DB_VOLUME=$(shell pwd)/db-data

.PHONY: api api-install api-update db db-follow db-stop db-remove gen-types fe fe-install fe-update install update

# ---------- Backend ----------
api:
	cd ./api && air

api-install:
	cd ./api && go mod tidy

api-update:
	cd ./api && go get -u all && go mod tidy

# ---------- Database ----------
db:
	@if [ $$(docker ps -a -q -f name=$(DB_NAME)) ]; then \
		echo "Starting existing database container..."; \
		docker start $(DB_NAME); \
	else \
		echo "Running database container for the first time..."; \
		docker run --name $(DB_NAME) -p $(DB_PORT):5432 -e POSTGRES_PASSWORD=$(DB_PASSWORD) -v $(DB_VOLUME):/var/lib/postgresql/data -d $(DB_IMAGE) -c max_connections=200; \
	fi

db-follow:
	docker logs $(DB_NAME) --follow
	
db-stop:
	echo "Stopping database container..."; \
	docker stop $(DB_NAME); \

db-remove:
	echo "Removing database container..."; \
	docker stop $(DB_NAME); \
	docker rm -f $(DB_NAME); \

gen-types:
	openapi-generator generate \
		-i ./api/generated/swagger.yaml \
		-g typescript-fetch \
		-o ./shared-fe/api-client/src/generated \
		--skip-validate-spec \
		--additional-properties=supportsES6=true,npmVersion=$$(npm --version),typescriptThreePlus=true
	cd ./shared-fe/api-client && yarn && yarn build

# ---------- Frontend ----------
fe:
	cd ./fe && yarn dev

fe-install:
	cd ./fe && yarn

fe-update:
	cd ./fe && yarn upgrade --latest

# ---------- Combined Targets ----------
install: api-install fe-install gen-types

update: api-update fe-update gen-types
