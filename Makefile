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
	./bin/$(CLI_NAME)

.PHONY: help
help:
	./bin/$(CLI_NAME) --help

.PHONY: test
test:
	go test -v -cover -coverprofile=index.out ./...

.PHONY: cover
cover: test
	go tool cover -html=index.out -o index.html
	python3 -m http.server 8765

.PHONY: deadcode
deadcode:
	deadcode ./...

.PHONY: govulncheck
govulncheck:
	govulncheck ./...

.PHONY: ci
ci: test deadcode govulncheck
