install:
	go get -t -d -v ./...

setup:
	@docker-compose up -d
	@mysql -h 127.0.0.1 -P 3306 -u root -e "CREATE DATABASE IF NOT EXISTS photon_development;"
	@docker run -it --network host bitnami/cassandra:latest cqlsh --username cassandra --password cassandra -e "create keyspace photon_development with replication={'class': 'SimpleStrategy', 'replication_factor': 1};"
	@make build
	-PHOTON_ENV=development bin/migrate up

build:
	go build -o ./bin/server ./cmd/server/main.go
	go build -o ./bin/migrate ./cmd/migrate/main.go

test:
	PHOTON_ENV=test go test -race -ldflags -s ./...

compile_protos:
	rm -rf protos && protoc -I ../ --go_out=plugins=grpc:. ../protos/*.proto

