package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-code/server_restful_json/person"
)

// rotas: localhost:8080            -> indice
//        localhost:8080/persons    -> lista pessoas
//        localhost:8080/persons/1  -> pessoa com id 1

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/persons", PersonIndex)
	router.HandleFunc("/persons/{personId}", PersonShow)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Bem-vindo!</h1>")
}

func PersonIndex(w http.ResponseWriter, r *http.Request) {
	persons := person.Persons{
		person.Person{Name: "Paulo C", Active: true, Create_at: time.Now()},
		person.Person{Name: "José da Silva"},
	}

	json.NewEncoder(w).Encode(persons)
}

func PersonShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["personId"]
	fmt.Fprintln(w, "<h1>Pessoa:", personId, "</h1>")
}
