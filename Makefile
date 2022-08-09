include dev.env

helloworld: 
	echo ${DB_NAME}

createdb:
	docker exec -it postgres-alpine14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-alpine14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

