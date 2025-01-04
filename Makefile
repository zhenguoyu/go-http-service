.PHONY: build test fmt clean run

build:
	go build -o bin/server ./cmd/httpserver
	go build -o bin/cli ./cmd/cli

test:
	go test -v ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/*
	go clean

run:
	go run ./cmd/httpserver
