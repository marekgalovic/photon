install_minikube:
	brew cask install minikube
	brew install kubectl

start_minikube:
	minikube start --vm-driver=xhyve

stop_minikube:
	minikube stop

compile_protos:
	rm -rf ./go/core/protos && mkdir ./go/core/protos && protoc -I ./protos/ --go_out=plugins=grpc:./go/core/protos ./protos/*.proto
	rm -rf ./ruby/client/lib/photon/protos && mkdir ./ruby/client/lib/photon/protos && protoc -I ./protos/ --ruby_out=plugins=grpc:./ruby/client/lib/photon/protos ./protos/*.proto
	rm -rf ./python/client/photon/protos && mkdir ./python/client/photon/protos && touch ./python/client/photon/protos/__init__.py && protoc -I ./protos/ --python_out=plugins=grpc:./python/client/photon/protos ./protos/*.proto

build:
	go build -o ./bin/server ./go/cmd/server/main.go
	go build -o ./bin/migrate ./go/cmd/migrate/main.go
	go build -o ./bin/deployer ./go/cmd/deployer/main.go
