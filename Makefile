makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=harrison40 -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root coding_challenge

dropdb:
	docker exec -it postgres12 dropdb coding_challenge

migrateup:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose down

loaddata:
	docker exec -i postgres12 psql -f - coding_challenge -t < .\db\migration\load_data.sql

generate:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
	
.PHONY: postgres createdb dropdb migrateup migratedown loaddata generate test server