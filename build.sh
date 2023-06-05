#!/bin/bash

rm -rf bin/
GOOS=darwin GOARCH=arm64 go build -o bin/GameOfLife-mac-arm main.go
GOOS=darwin GOARCH=amd64 go build -o bin/GameOfLife-mac-x64 main.go
GOOS=windows GOARCH=amd64 go build -o bin/GameOfLife-win-x64 main.go
GOOS=linux GOARCH=amd64 go build -o bin/GameOfLife-linux-x64 main.go
