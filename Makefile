migration:
	goose postgres "user=fintechAdmin password=12345 dbname=fintechDB host=localhost port=5432 sslmode=disable" up
inMemory:
	docker compose -f inMemory.yml up -d
postgres:
	docker compose -f postgres.yml up -d
