#!/usr/bin/env bash
set +e

echo "Initializing Go modules.."
GO111MODULE=on go mod init $1
