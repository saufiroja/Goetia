.PHONY run server dev:
dev:
	@echo "Running server..."
	@cd cmd && go run main.go

.PHONY up docker dev:
docker-up:
	@echo "Running docker..."
	@docker-compose -f docker-compose-dev.yaml up -d

.PHONY run unit test:
unit-test:
	@echo "Running unit tests..."
	@cd test/unit_test && go test -v ./...

.PHONY run integration test:
integration-test:
	@echo "Running integration tests..."
	@cd test/integration_test && go test -v ./...

.PHONY run e2e test:
e2e-test:
	@echo "Running e2e tests..."
	@cd test/e2e_test && go test -v ./...

.PHONY run all test:
all-test:
	@echo "Running all tests..."
	@cd test && go test -v ./...

.PHONY down docker dev:
docker-down:
	@echo "Stopping docker..."
	@docker-compose -f docker-compose-dev.yaml down -v