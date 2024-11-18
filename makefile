create-network:
	docker network create foodCatalog

mongo:
	docker run -d --network foodCatalog --name foodCatalog \
		-p 27017:27017 \
		mongo

mongo-rm:
	docker stop foodCatalog
	docker rm foodCatalog

serve:
	go run main.go

tidy:
	go mod tidy