package route

import (
	"go-code/server_restful_json/v2/controller"
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
// [GET]  localhost:8080           handlers.Index        -> indice(welcome)
// [GET]  localhost:8080/persons   handlers.PersonIndex  -> lista pessoas em array de jsons
// [GET]  localhost:8080/persons/1 handlers.PersonShow   -> pessoa com id 1 em json
// [POST] localhost:8080/persons   handlers.PersonCreate -> cadastro de pessoa via json

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"UsuarioIndex",
		"GET",
		"/usuarios",
		controller.UsuarioIndex,
	},
	Route{
		"UsuarioShow",
		"GET",
		"/usuarios/{usuarioId}",
		controller.UsuarioShow,
	},
	Route{
		"UsuarioCreate",
		"POST",
		"/usuarios",
		controller.UsuarioCreate,
	},
	Route{
		"UsuarioRemove",
		"DELETE",
		"/usuarios/{usuarioId}",
		controller.UsuarioRemove,
	},
	Route{
		"UsuarioAlter",
		"PUT",
		"/usuarios",
		controller.UsuarioAlter,
	},
}
