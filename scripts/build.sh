#!/bin/bash

gox -output "../dist/dioderAPI_{{.OS}}_{{.Arch}}" -parallel=2 -verbose -os="${OPERATING_SYSTEM}" -ldflags="-s -w"
