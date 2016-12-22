#!/bin/bash

export PATH="$PATH:$PWD/bin"
export GOPATH="$PWD"

export http_proxy=http://172.16.222.3:9980/
export https_proxy=http://172.16.222.3:9980/
export socket_proxy=http://172.16.222.3:9980/

go get -v ./...
go get -v ./...

# Build application
go build ./...

#go install -v -gcflags "-N -l" ./...
