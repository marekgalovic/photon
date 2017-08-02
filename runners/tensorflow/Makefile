install_tf:
	export TF_TYPE="cpu"
	export TARGET_DIRECTORY="/usr/local"
	curl -L \
   		"https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-$(go env GOOS)-x86_64-1.2.1.tar.gz" | sudo tar -C $(TARGET_DIRECTORY) -xz

install:
	go get -t -d -v ./...

build:
	go build -o ./bin/photon_tensorflow_runner ./cmd/runner/main.go

test:
	PHOTON_ENV=test go test -ldflags -s ./...

compile_protos:
	rm -rf protos && protoc -I ../ --go_out=plugins=grpc:. ../protos/runner/*.proto

