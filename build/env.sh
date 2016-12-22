#!/bin/bash

set -x

export http_proxy=http://172.16.222.3:9980/
export https_proxy=http://172.16.222.3:9980/
export socket_proxy=http://172.16.222.3:9980/

# Set up environment
export PATH="$PATH:/var/lib/jenkins/go/bin"

echo $PATH

# BASE.LIB
export GOPATH="$PWD"

# LIB
go get -v github.com/nsf/gocode
go get -v golang.org/x/tools/cmd/guru
go get -v github.com/rogpeppe/godef
go get -v github.com/tebeka/go2xunit
go get -v github.com/stretchr/testify/assert
go get -v github.com/axw/gocov/gocov
go get -v gopkg.in/matm/v1/gocov-html
go get -v github.com/t-yuki/gocov-xml
go get -v github.com/golang/lint/golint
#go install -v -gcflags "-N -l" ./...
