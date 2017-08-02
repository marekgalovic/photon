install:
	go get -t -d -v ./...

build:
	go build -o ./bin/photon_tensorflow_runner ./cmd/runner/main.go

test:
	PHOTON_ENV=test go test -ldflags -s ./...

compile_protos:
	rm -rf protos && protoc -I ../ --go_out=plugins=grpc:. ../protos/runner/*.proto

