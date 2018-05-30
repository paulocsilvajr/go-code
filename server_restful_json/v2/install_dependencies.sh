#!/bin/bash

install(){
    echo "$1:"
    go get -v $1
}

echo -e "Instalando dependências\n"

install github.com/gorilla/mux
install github.com/mattn/go-sqlite3
