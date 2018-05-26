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
