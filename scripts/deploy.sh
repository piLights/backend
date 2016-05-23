#!/bin//bash

/bin/bash scripts/build.sh

if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
	ghr --username piLights --token $GITHUB_TOKEN --replace --prerelease --debug pre-release dist/
fi
