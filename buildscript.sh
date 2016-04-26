#!/bin/bash

if [ ! -d "/tmp/CIgo" ]; then
  mkdir /tmp/CIgo
fi

cd src/

# Get all dependencies
go get

cd ..
mkdir builds

go build -o builds/dioderAPI_x86 -v src/*.go
env GOOS=linux GOARCH=arm go build -o builds/dioderAPI_arm -v src/*.go
