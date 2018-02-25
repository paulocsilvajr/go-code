package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-code/server_restful_json/person"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Bem-vindo!</h1>")
}

func PersonIndex(w http.ResponseWriter, r *http.Request) {
	persons := person.Persons{
		person.Person{Name: "Paulo C", Active: true, Create_at: time.Now()},
		person.Person{Name: "José da Silva"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(persons); err != nil {
		panic(err)
	}
}

func PersonShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["personId"]
	fmt.Fprintln(w, "<h1>Pessoa:", personId, "</h1>")
}
