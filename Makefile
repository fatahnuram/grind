help:
	@echo "possible commands: help, install, run, build, test"

install:
	go install

run:
	go run .

build:
	go build

test:
	go test -v ./...

.PHONY: help install run build test
