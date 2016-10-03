#!/bin/bash

for file in $(find dist -type f); do
	HASH=$(sha256sum $file | awk '{print $1}')
	echo $HASH > $file.sha256
done
