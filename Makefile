.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build

.PHONY: run
run:
	go run ./cmd/apiserver