.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := go-strcmp
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS  := -X 'main.revision=$(REVISION)'

## Setup the development environment
.PHONY: setup
setup:
	go get -u github.com/rakyll/gotest
	go get -u golang.org/x/lint/golint

## Cleanup unnecessary modules, etc.
.PHONY: cleanup
cleanup:
	go mod tidy

## Run tests
.PHONY: test
test:
	gotest ./... -v -coverprofile=c.out

## Run benchmark tests
.PHONY: bench
bench:
	gotest ./... -bench=.

## Run lint
.PHONY: lint
lint:
	golint -set_exit_status ./...
