test: config mock
	go test -gcflags=-l -v ./... -coverprofile cover.out
	go tool cover -func cover.out

mock:
	mockgen -destination=mocks/redis_mock.go -package=mocks -source=internal/redis/provider.go

config:
	cp example.config.yaml config.yaml
