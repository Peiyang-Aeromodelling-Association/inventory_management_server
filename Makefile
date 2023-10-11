PROJECTNAME=$(shell basename "$(PWD)")

# if found secret.env file, include it.
ifneq ("$(wildcard secret.env)","")
	include secret.env
endif

DB_NAME=$(PROJECTNAME)_db
DB_URL=postgresql://postgres:$(POSTGRES_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable

## postgres: Run postgres container
postgres:
	@echo "Running postgres container"
	docker run --name postgres --restart=always -p 5432:5432 -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_USER=postgres -d postgres:alpine

## createdb: Create database
createdb:
	@echo "Creating database with user postgres"
	docker exec -it postgres createdb --username=postgres --owner=postgres $(DB_NAME)

## dropdb: Drop database
dropdb:
	@echo "Dropping database '$(DB_NAME)'"
	docker exec -it postgres dropdb --username=postgres $(DB_NAME)

## migrateup: Migrate database up
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

## migratedown: Migrate database down
migratedown:
	echo 'y' | migrate -path db/migration -database "$(DB_URL)" -verbose down

## sqlc: Generate sqlc
sqlc:
	sqlc generate

## test: Run test
test:
	go test -v -cover ./...

## wipedb: Drop tables and create tables
wipedb: migratedown migrateup

.PHONY: help postgres createdb dropdb migrateup migratedown sqlc test dbclean

.DEFAULT_GOAL := help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in '"$(PROJECTNAME)"':"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'