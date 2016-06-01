#!/bin/bash

cd src/

gox -output -ldflags="-s -w" "../dist/dioderAPI_{{.OS}}_{{.Arch}}" -parallel=2 -verbose -os="${OPERATING_SYSTEM}"

cd ..
