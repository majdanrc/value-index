.PHONY: all build run test

BINARY_NAME=value-indexer

all: build

build:
	go build -o $(BINARY_NAME) ./cmd/server

run: build
	./$(BINARY_NAME)

test:
	go test ./... -v

clean:
	rm -f $(BINARY_NAME)