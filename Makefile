SHELL := /bin/bash

.DEFAULT_GOAL := build

.PHONY: clean cleanup init test build

clean:
	@rm -Rf output

init: clean
	@mkdir output
	@mkdir output/bin
	@mkdir output/tools
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b output/tools v1.38.0
	@output/tools/golangci-lint --version

	go get golang.org/x/tools/cmd/goimports

deps: init
	go get && go mod tidy

cleanup:
	gofmt -w .
	goimports -w .

linter: deps cleanup
	@output/tools/golangci-lint run --enable-all --deadline 5m0s --out-format checkstyle | tee ./output/linter.out

test: init cleanup
	go test ./... -v -race -covermode=atomic -coverprofile ./output/coverage.out -json | tee ./output/test_result.json

build: test
	env GOOS=linux GOARCH=amd64 go build -o output/bin/
	
upgrade:
	go get -u
	