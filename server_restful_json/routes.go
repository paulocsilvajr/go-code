package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
