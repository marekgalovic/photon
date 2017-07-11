install:
	go get -t -d -v ./...

compile_protos:
	rm -rf protos && protoc -I ../ --go_out=plugins=grpc:. ../protos/*.proto
