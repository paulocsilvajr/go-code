package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	fmt.Fprintln(w, "<h1>Pessoa index</h1>")
}

func PersonShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["personId"]
	fmt.Fprintln(w, "<h1>Pessoa:", personId, "</h1>")
}
