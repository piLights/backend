#!/bin//bash

cd src/
gox -output "../dist/dioderAPI_{{.OS}}_{{.Arch}}" -parallel=2 -verbose -os="${OPERATING_SYSTEM}"
cd ..

if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
	ghr --username piLights --token $GITHUB_TOKEN --replace --prerelease --debug pre-release dist/
fi
