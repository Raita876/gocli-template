VERSION := $(shell cat ./VERSION)
CLI_NAME := hello

.PHONY: build
build:
	go build \
		-o ./bin/$(CLI_NAME) \
		-ldflags "-X main.version=$(VERSION) -X main.name=$(CLI_NAME)" \
		./cmd/hello

.PHONY: run
run: build
	./bin/$(CLI_NAME) --verbose --exclnum 10 "World"

.PHONY: help
help: build
	./bin/$(CLI_NAME) --help

.PHONY: version
version: build
	./bin/$(CLI_NAME) -V

.PHONY: test
test:
	go test -v -cover -coverprofile=index.out ./...

.PHONY: cover
cover: test
	go tool cover -html=index.out -o index.html
	python3 -m http.server 8765

.PHONY: golangci
golangci:
	golangci-lint run -v ./...

.PHONY: govulncheck
govulncheck:
	govulncheck ./...

.PHONY: ci
ci: test golangci govulncheck
