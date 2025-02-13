.PHONY: test lint build run run-dev format install-dependencies

format:
	gofmt -d -w .

vet:
	go vet ./...

lint:
	./bin/golangci-lint run ./...

build:
	go build -o ./build/server ./cmd/server.go

run: build
	./build/server

run-dev:
	go run ./cmd/server.go

install-dependencies:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.63.4

test:
	go test -v ./...

ci: install-dependencies format vet lint test