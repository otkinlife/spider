#!/bin/sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o output/mac/spider $1
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o output/win/spider.bat $1
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/linux/spider $1