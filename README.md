## When using

- Copy the structure
- Replace every `project-template` with the project template name

## API

- `cp .env.example .env` and edit it.
- `air`

## Database

```
# First time:
docker run --name project-template_db -p 42376:5432 -e POSTGRES_PASSWORD=s0m3C0mpl3xP4ss -v $(pwd)/db-data:/var/lib/postgresql/data -d postgres:17.2-alpine3.20 -c max_connections=200
docker container logs project-template_db --follow

# Every other time:
docker start project-template_db
docker container logs project-template_db --follow
```
