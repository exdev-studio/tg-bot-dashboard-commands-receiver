.PHONY: clean build

clean:
	rm -rf bin

build:
	go build -ldflags "-s -w" -o bin/receiver ./cmd/receiver

.DEFAULT_GOAL := build
