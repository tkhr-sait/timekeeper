#!/bin/bash
cd $(dirname $0)
export GOPATH=$(cd ../;pwd)
env GOOS=darwin GOARCH=amd64 go build -o ../bin/timekeeper.mac ../src/*.go
env GOOS=windows GOARCH=amd64 go build -o ../bin/timekeeper.exe ../src/*.go
cd -
