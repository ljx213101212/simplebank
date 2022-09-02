include app.env
include prod.env

createdb:
	docker exec -it postgres-alpine14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-alpine14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migrateuprds:
	migrate -path db/migration -database "${DB_SOURCE_RDS}" -verbose up

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

dockerbuild:
	docker build -t simplebank:latest .

dockerrun:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres-alpine14:5432/simple_bank?sslmode=disable" simplebank:latest

dockerstop:
	docker container stop simplebank

dockerrm:
	docker container rm simplebank

dockerippostgres:
	docker container inspect postgres-alpine14 | grep IPAddress

dockeripsimplebank:
	docker container inspect simplebank | grep IPAddress

dockercreatenetwork:
	docker network create bank-network

dockerinspectnetwork:
	docker network inspect bank-network

dockerconnectpostgres:
	docker network connect bank-network postgres-alpine14

dockerconnectsimplebank:
	docker network connect bank-network simplebank

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock
