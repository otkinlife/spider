#!/bin/sh
cd ..
go build -o output/mac/spider tool.go
go build -o output/mac/client client.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o output/win/spider.bat tool.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o output/win/client.bat client.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/linux/spider tool.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/linux/client client.go