install:
	go get -u cloud.google.com/go/...
	go get -t -d -v ./...

generate_mocks:
	# Repositories
	mockgen --source repositories/credentials_repository.go --destination repositories/mock/credentials_repository_mock.go --imports .=github.com/marekgalovic/photon/go/core/repositories
	mockgen --source repositories/deployer_repository.go --destination repositories/mock/deployer_repository_mock.go
	mockgen --source repositories/features_repository.go --destination repositories/mock/features_repository_mock.go --imports .=github.com/marekgalovic/photon/go/core/repositories
	mockgen --source repositories/instances_repository.go --destination repositories/mock/instances_repository_mock.go --imports .=github.com/marekgalovic/photon/go/core/repositories
	mockgen --source repositories/models_repository.go --destination repositories/mock/models_repository_mock.go --imports .=github.com/marekgalovic/photon/go/core/repositories
	# Storage
	mockgen --source storage/features/features_store.go --destination storage/features/mock/features_store_mock.go
	mockgen --source storage/files/files_store.go --destination storage/files/mock/files_store_mock.go
	# Protos
	mkdir -p protos/mock
	mockgen --source protos/evaluator.pb.go --destination protos/mock/evaluator_mock.pb.go --imports .=github.com/marekgalovic/photon/go/core/protos
	mockgen --source protos/runner.pb.go --destination protos/mock/runner_mock.pb.go --imports .=github.com/marekgalovic/photon/go/core/protos


test:
	PHOTON_ENV=test go test -ldflags -s ./...
