#!/bin/bash

for file in $(find ../dist -type f -name "dioderAPI*"); do
	curl -Iv --user "${USERNAME}:${PASSWORD}" -T $file https://binupload.pilights.jf-projects.de
done
