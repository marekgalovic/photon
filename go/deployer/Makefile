install:
	go get -t -d -v ./...

build:
	go build -o ./bin/deployer ./cmd/deployer/main.go

test:
	PHOTON_ENV=test go test -ldflags -s ./...
