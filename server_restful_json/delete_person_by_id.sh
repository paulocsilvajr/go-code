if [ -z $1 ]; then
    echo "Informe um ID"
else
    curl -X DELETE http://localhost:8080/persons/$1
fi
