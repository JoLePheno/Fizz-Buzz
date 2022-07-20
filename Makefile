.PHONY: test integration

GO:=go

# Run all test
test:
	LOG_FORMAT=console $(GO) test -mod=vendor -count=1 -race -cover -short ./...

# Run integration test in docker container
docker-integration:
	docker-compose build
	docker-compose run app make integration
	docker-compose down

integration:
	LOG_FORMAT=console $(GO) test -mod=vendor -count=1 -race -cover -run TestIntegration ./...