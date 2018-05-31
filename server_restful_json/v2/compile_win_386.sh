#!/bin/bash

PASTA=bin/win386
ARQUIVO=$PASTA/server_rest.exe

export GOOS=windows
export GOARCH=386
export CGO_ENABLED=0

mkdir $PASTA 2> /dev/null
go build -o $ARQUIVO -a -ldflags '-extldflags "-static"' && echo "Gerado executável: $PASTA/$(ls $PASTA)"
