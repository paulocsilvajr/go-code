package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// rotas:
// [GET]    localhost:8080           handlers.Index         -> indice(welcome)
// [GET]    localhost:8080/persons   handlers.PersonIndex   -> lista pessoas em json array
// [GET]    localhost:8080/persons/1 handlers.PersonShow    -> pessoa com id 1 em json
// [POST]   localhost:8080/persons   handlers.PersonCreate  -> cadastro de pessoa via json
// [DELETE] localhost:8080/persons/1 handlers.PersonDestroy -> remove pessoa com id 1
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
	Route{
		"PersonCreate",
		"POST",
		"/persons",
		PersonCreate,
	},
	Route{
		"PersonDestroy",
		"DELETE",
		"/persons/{personId}",
		PersonDestroy,
	},
}
