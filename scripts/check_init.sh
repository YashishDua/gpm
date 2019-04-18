#!/usr/bin/env bash

file=".goboil"

if [ -f "$file" ]
then
	echo "$file already present."
else
	GO111MODULE=on go mod init $1
fi
