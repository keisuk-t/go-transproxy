#!/bin/sh

VERSION=0.5.1

go build -v -o bin/transproxy cmd/transproxy/main.go
# tar cvzf go-transproxy-$VERSION.tar.gz bin README.md LICENSE
