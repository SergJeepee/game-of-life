#!/bin/bash

rm -rf bin/
GOOS=darwin GOARCH=arm64 go build -o bin/GameOfLife-mac-arm main.go
GOOS=darwin GOARCH=amd64 go build -o bin/GameOfLife-mac-int main.go
GOOS=windows GOARCH=amd64 go build -o bin/GameOfLife-win main.go
GOOS=linux GOARCH=amd64 go build -o bin/GameOfLife-linux main.go
