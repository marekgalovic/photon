install:
	go get -t -d -v ./...

generate_mocks:
	mockgen --source evaluator/features_resolver.go --destination evaluator/mock/features_resolver_mock.go
	mockgen --source evaluator/model_resolver.go --destination evaluator/mock/model_resolver_mock.go

test:
	PHOTON_ENV=test go test -ldflags -s ./...
