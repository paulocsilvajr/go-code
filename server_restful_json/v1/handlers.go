package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-code/server_restful_json/v1/person"
)

// funções acessadas pelas rotas(routes.go)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Bem-vindo!</h1>")
}

func PersonIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(person.ListPersons); err != nil {
		panic(err)
	}
}

func PersonShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["personId"]

	id, err := strconv.Atoi(personId)
	if err != nil {
		panic(err)
	}

	found_person := person.DaoFindPerson(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if found_person.Id != 0 {
		if err := json.NewEncoder(w).Encode(found_person); err != nil {
			panic(err)
		}
	}
}

func PersonCreate(w http.ResponseWriter, r *http.Request) {
	var new_person person.Person

	// io.LimitReader define limite para o tamanho do json
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &new_person); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := person.DaoCreatePerson(new_person)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
