create-network:
	docker network create food-catalog

pg:
	docker run -d \
		--name food-catalog \
		--network food-catalog \
		-p 5432:5432 \
  	-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
  	postgres

pgrm:
	docker stop food-catalog
	docker rm food-catalog

createdb:
	docker exec -it food-catalog createdb --username=root --owner=root food-catalog

dropdb:
	docker exec -it food-catalog dropdb food-catalog

schema:
	migrate create -ext sql -dir external/database/migrations -seq init_schema

migrateup:
	migrate -path database/postgres/migrations -database "postgresql://root:secret@localhost:5432/food-catalog?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migrations -database "postgresql://root:secret@localhost:5432/food-catalog?sslmode=disable" -verbose down

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