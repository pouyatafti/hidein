#!/bin/sh

mkdir -p bin
go build -o bin/encode github.com/pouyatafti/hidein/cmd/encode
go build -o bin/decode github.com/pouyatafti/hidein/cmd/decode
