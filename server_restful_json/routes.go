package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// rotas: localhost:8080            -> indice
//        localhost:8080/persons    -> lista pessoas
//        localhost:8080/persons/1  -> pessoa com id 1

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PersonIndex",
		"GET",
		"/persons",
		PersonIndex,
	},
	Route{
		"PersonShow",
		"GET",
		"/persons/{personId}",
		PersonShow,
	},
}
