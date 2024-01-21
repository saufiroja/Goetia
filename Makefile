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

.PHONY generate proto:
protoc:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@echo "Generating protobuf files..."
	@protoc -I ./proto --go_out=./ \
			--go-grpc_out=require_unimplemented_servers=false:./ \
			--grpc-gateway_out . --grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt generate_unbound_methods=true \
			./proto/todos/*.proto ./proto/google/*.proto
	@echo "Done"