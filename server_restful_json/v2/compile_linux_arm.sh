#!/bin/bash

mkdir bin/arm
env GOOS=linux GOARCH=arm go build -o bin/arm/server_rest
