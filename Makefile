.PHONY: build clean run fmt

GOFLAGS ?= $(GOFLAGS:)

build:
	@go build $(GOFLAGS) -o builds/dioderAPI *.go

clean:
	@go clean

run:
	@go run *.go