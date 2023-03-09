postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root carpool

enterdb:
	docker exec -it postgres12 psql -U root

dropdb:
	docker exec -it postgres12 dropdb carpool

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/carpool?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/carpool?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server