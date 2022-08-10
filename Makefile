include app.env

createdb:
	docker exec -it postgres-alpine14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-alpine14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -destination db/mock/store.go github.com/ljx213101212/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock

