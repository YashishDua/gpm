#!/usr/bin/env bash

echo "Initializing Go modules.."
GO111MODULE=on go mod init $1
