install:
	go get -t -d -v ./...

test:
	PHOTON_ENV=test go test -ldflags -s ./...

compile_protos:
	rm -rf protos && protoc -I ../../ --go_out=plugins=grpc:. ../../protos/*.proto	
