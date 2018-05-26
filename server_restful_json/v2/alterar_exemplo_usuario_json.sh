#!/bin/bash

curl -X "PUT" -H "Content-Type: application/json" -d '{"id": 4, "nome":"Vergil", "senha":"Sparda"}' http://localhost:8080/usuarios
