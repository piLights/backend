.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

build:
	@go build $(GOFLAGS) -o builds/dioderAPI src/*.go

clean:
	@go clean

proto:
	@protoc --go_out=plugins=grpc:. src/proto/*.proto

run:
	@go run src/*.go
