#!/bin/bash

mkdir bin/win386
env GOOS=windows GOARCH=386 go build -o bin/win386/server_rest.exe
