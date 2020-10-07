#!/bin/bash
set -e
cd src
echo "build in progress ..."
echo "GOPATH : " $GOPATH
echo "GOBIN : " $GOBIN
go mod download
go get -d -v
GO111MODULE=on CGO_ENABLED=0 GOOS=windows go build -o git-pull .

echo "build terminated, application build in $GOBIN path."
