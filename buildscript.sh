#!/bin/bash

if [ ! -d "/tmp/CIgo" ]; then
  mkdir /tmp/CIgo
fi


# Get all dependencies
go get

mkdir builds

# Replace the version-string
#If exists
if [ ! -z ${CI_BUILD_REF+x} ]; then
	find . -type f -name "*.go" | xargs sed -i -e "s/debugVersion/$CI_BUILD_REF/g"
fi

if [ ! -z ${TRAVIS_COMMIT+x} ]; then
    find . -type f -name "*.go" | xargs sed -i -e "s/debugVersion/$TRAVIS_COMMIT/g"
fi

go build -o builds/dioderAPI_x86 -v *.go
env GOOS=linux GOARCH=arm go build -o builds/dioderAPI_arm -v *.go
