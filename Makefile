DB_NAME=project-template_db
DB_PORT=35232
DB_PASSWORD=s0m3C0mpl3xP4ss
DB_IMAGE=postgres:17.2-alpine3.20
DB_VOLUME=$(shell pwd)/db-data

.PHONY: db gen-types fe fe-install

# ---------- Backend ----------
api:
	cd ./api && air

db:
	@if [ $$(docker ps -a -q -f name=$(DB_NAME)) ]; then \
		echo "Starting existing database container..."; \
		docker start $(DB_NAME); \
	else \
		echo "Running database container for the first time..."; \
		docker run --name $(DB_NAME) -p $(DB_PORT):5432 -e POSTGRES_PASSWORD=$(DB_PASSWORD) -v $(DB_VOLUME):/var/lib/postgresql/data -d $(DB_IMAGE) -c max_connections=200; \
	fi
	docker logs $(DB_NAME) --follow

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