install:
	go get -t -d -v ./...

build:
	go build -o ./bin/server ./cmd/server/main.go
	go build -o ./bin/migrate ./cmd/migrate/main.go

test:
	PHOTON_ENV=test go test -ldflags -s ./...

compile_protos:
	rm -rf protos && protoc -I ../ --go_out=plugins=grpc:. ../protos/*.proto

