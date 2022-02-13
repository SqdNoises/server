postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

redis:
	docker run -d --name redis -p 6379:6379  -v /path/to/redisconf/redis.conf:/redis.conf redis redis-server /redis.conf

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root test

dropdb:
	docker exec -it postgres12 dropdb test

recreate:
	make dropdb && make createdb

start:
	docker start postgres12 && docker start redis

test:
	go test -v -cover ./service/... ./handler/...

e2e:
	go test github.com/aelpxy/xoniaapp

lint:
	golangci-lint run

mock:
	mockery --all

build:
	go github.com/aelpxy/xoniaapp

fmt:
	go fmt github.com/aelpxy/...

swag:
	swag init

workflow:
	make fmt && make lint && make test
