# unit-test-example
Golang examples with write unit test cases

## Makefile
- `make build` to build the golang binary of application
- `make run` to run application on local
- `make clean` to clean generated binary of application

## Run on local
- `make run`

## Unit Testing
- `go test -short ./...` to run without integration tests
- `go test -v ./...` to run with integration tests

## Health Check API
- `http://localhost:8080/health`