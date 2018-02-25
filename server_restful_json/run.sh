#!/bin/bash

if [ "$1" == "-h" ]; then
    echo "Sintaxe: $0 [ -t | -j ]"
    echo "         -t testes(go run)"
    echo "         -h Ajuda"
    echo "         Quando omite-se o parâmetro, faz o build(go build) e executa"

elif [ "$1" == "-t" ]; then
    go run *.go

else
    ./build.sh
    ./server_restful_json

fi
