#!/bin/bash

export PATH="$PATH:$PWD/bin"
export GOPATH="$PWD"

export http_proxy=http://172.16.222.3:9980/
export https_proxy=http://172.16.222.3:9980/
export socket_proxy=http://172.16.222.3:9980/

go get -v ./...
go get -v ./...

# Run tests (JUnit plugin)
rm -f test.out
touch tmp.out
echo "mode: set" > coverage.out
for pkg in $(go list ./...);
do
    if [[ $pkg != *"github.com"* || $pkg != *"gopkg.in"* ]]; then
      echo "testing... ${pkg}"
      go test -v -coverprofile=tmp.out ${pkg} >> test.out
      if [ -f tmp.out ]; then
         cat tmp.out | grep -v "mode: set" >> coverage.out
      fi
    fi  
done

#go test vendor/app/utilities -coverprofile=tmp.out ${pkg} >> test.out
rm -f ./tmp.out
cat test.out | go2xunit -output tests.xml

# Generate coverage reports (Cobertura plugin)
gocov convert coverage.out | gocov-xml > cobertura-coverage.xml
gocov convert coverage.out | gocov-html > cobertura-coverage.html

# Run lint tools (Compiler warning plugin)
golint ./... > lint.txt

# Run vet tools (Compiler warning plugin)
go vet ./... > vet.txt

# Build application
go build ./...

#go install -v -gcflags "-N -l" ./...
