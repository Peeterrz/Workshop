@echo off

set GOPATH=%CD%
go get -v github.com/nsf/gocode
go get -v golang.org/x/tools/cmd/guru
go get -v github.com/rogpeppe/godef
go get -v github.com/yvasiyarov/swagger
go get -v golang.org/x/crypto/bcrypt
go get -v gopkg.in/mgo.v2
go get -v gopkg.in/mgo.v2/bson
go get -v ./...
go get -v ./...

pause