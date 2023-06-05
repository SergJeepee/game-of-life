#!/bin/bash

GOOS=darwin GOARCH=arm64 go build -o bin/GameOfLife-mac-arm GameOfLife.go
GOOS=darwin GOARCH=amd64 go build -o bin/GameOfLife-mac-int GameOfLife.go
GOOS=windows GOARCH=amd64 go build -o bin/GameOfLife-win GameOfLife.go
GOOS=linux GOARCH=amd64 go build -o bin/GameOfLife-linux GameOfLife.go
