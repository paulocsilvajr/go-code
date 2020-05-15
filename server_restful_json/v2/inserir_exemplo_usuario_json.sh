#!/bin/bash

curl -H "Content-Type: application/json" -d '{"nome":"Dante", "senha":"ABC"}' http://localhost:8080/usuarios
