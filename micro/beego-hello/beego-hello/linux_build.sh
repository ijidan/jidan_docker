#!/bin/bash

export CGO_ENABLED=0
export GOARCH=amd64
export GOOS=linux

go build -a -v -o beego-hello main.go

tar -czvf beego-hello.tar.gz conf/ views/ beego-hello