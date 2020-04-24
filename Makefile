# .PHONY: lint test
## Run lint, test
all:
	make lint
	make test

## Run golangci-lint using .golangci.yml config
lint:
	golangci-lint run --config .golangci.yml

## Run start_images, - go test ./... -cover -race and then stop_images
test:
	- go test ./... -cover

## Generates mocks for all interfaces
generate-mocks:
	mockery -all -case underscore

## Generates models from swagger.yml specification.
generate-swagger-models:
	GO111MODULE=off swagger generate model --spec=swagger.yml
