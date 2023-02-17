.PHONY: test fmt build clean

all: fmt lint test build

build:
	go build -o build/wyag ./bin/wyag

test:
	go test -v ./...

lint:
	go vet ./...

fmt:
	go fmt ./...

clean:
	rm -rf build
