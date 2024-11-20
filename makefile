create-network:
	docker network create foodCatalog

pg:
	docker run -d \
		--name foodCatalog \
		--network foodCatalog \
		-p 5432:5432 \
  	-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
  	postgres

pgrm:
	docker stop foodCatalog
	docker rm foodCatalog

createdb:
	docker exec -it foodCatalog createdb --username=root --owner=root foodCatalog

dropdb:
	docker exec -it foodCatalog dropdb foodCatalog

schema:
	migrate create -ext sql -dir external/database/migrations -seq init_schema

migrateup:
	migrate -path database/postgres/migrations -database "postgresql://root:secret@localhost:5432/foodCatalog?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migrations -database "postgresql://root:secret@localhost:5432/foodCatalog?sslmode=disable" -verbose down

sqlc:
	sqlc generate

gqlgen:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.56
	go get github.com/99designs/gqlgen@v0.17.56
	go run github.com/99designs/gqlgen generate
	go mod tidy

tidy:
	go mod tidy

serve:
	go run cmd/main.go