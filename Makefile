network:
	docker network create db-network

postgres:
	docker run --name api-carpool-db -p 5432:5432 --network db-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it api-carpool-db createdb --username=root --owner=root carpool

enterdb:
	docker exec -it api-carpool-db psql -U root

dropdb:
	docker exec -it api-carpool-db dropdb carpool

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/carpool?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/carpool?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server-build:
	docker build --tag api-carpool .

server-run:
	docker run -d --name api-carpool --network db-network -p 8080:8080 api-carpool

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server