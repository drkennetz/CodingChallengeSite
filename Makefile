makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=harrison40 -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root coding_challenge

dropdb:
	docker exec -it postgres12 dropdb coding_challenge

migrateup:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:harrison40@localhost:5432/coding_challenge?sslmode=disable" --verbose down 1

loaddata:
	docker exec -i postgres12 psql -f - coding_challenge -t < .\db\migration\load_data.sql

loadtestdata:
	docker exec -i postgres12 psql -f - coding_challenge -t < ./db/migration/load_data.sql

generate:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/drkennetz/codingchallenge/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 loaddata loadtestdata generate test server mock