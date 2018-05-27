package controller

import (
	"encoding/json"
	"fmt"
	"go-code/server_restful_json/v2/model/usuario"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Bem-vindo</h1>")
}

func UsuarioIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	usuario.DaoCarregaUsuarios()

	if err := json.NewEncoder(w).Encode(usuario.ListaUsuarios); err != nil {
		panic(err)
	}
}

func UsuarioShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioId := vars["usuarioId"]

	id, err := strconv.Atoi(usuarioId)
	if err != nil {
		panic(err)
	}

	var usuarioEncontrado *usuario.Usuario
	usuarioEncontrado, err = usuario.DaoProcuraUsuario(id)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(*usuarioEncontrado); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
	}
}

func UsuarioCreate(w http.ResponseWriter, r *http.Request) {
	var novoUsuario usuario.Usuario

	// io.LimitReader define limite para o tamanho do json
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &novoUsuario); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u := usuario.DaoAdicionaUsuario(novoUsuario)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func UsuarioRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioId := vars["usuarioId"]

	id, err := strconv.Atoi(usuarioId)
	if err != nil {
		panic(err)
	}

	err = usuario.DaoRemoveUsuario(id)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
	}
}

func UsuarioAlter(w http.ResponseWriter, r *http.Request) {
	var novoUsuario usuario.Usuario

	// io.LimitReader define limite para o tamanho do json
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &novoUsuario); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u := usuario.DaoAlteraUsuario(novoUsuario)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}
