#!/usr/bin/env bash

echo "Initializing Go modules.."
file="go.mod"
if [ -f "$file" ]
then
	echo "$file already present."
else
	GO111MODULE=on go mod init $1
fi
