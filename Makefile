server:
	go run main.go
swagger:
	swagger generate spec -o swagger.json
postgres:
	run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres 
createdb:
	docker exec -it postgres createdb --username=root --owner=root bookstore
dropdb:
	docker exec -it postgres dropdb bookstore
migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/postgres?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
