# Smartmei Microservice

## Introduction

This service handles requests to get information about Smartmei website



## Configuration

### Instal `golang`

- https://golang.org/dl/

### Get dependencies

- `go get`

### Install `golangci-lint`

- https://github.com/golangci/golangci-lint

### Install `mockery`

- https://github.com/vektra/mockery

### Install `go-swagger`

- https://github.com/go-swagger/go-swagger/blob/master/docs/install.md

## Functionalities

- `make lint`: Runs linters to check code styles
- `make test`: Runs tests with coverage
- `make generate-mocks`: Generates mocks
- `make generate-swagger-models`: Generates models from swagger file

### Running project

- `go run main.go`: Will start on `localhost:5656`

### Apis

- `curl -XGET localhost:5656/smartmei/crawler?url=https://www.smartmei.com.br`: Gets information about transference price converted to `USD`, `EUR` and `BRL`
- `curl -XGET localhost:5656/smartmei/heartbeat`: Keep alive endpoint
