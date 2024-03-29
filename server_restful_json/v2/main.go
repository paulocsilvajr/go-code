package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulocsilvajr/go-code/server_restful_json/v2/config"
	"github.com/paulocsilvajr/go-code/server_restful_json/v2/config/route"
	"github.com/paulocsilvajr/go-code/server_restful_json/v2/helper"
	"github.com/paulocsilvajr/go-code/server_restful_json/v2/model/usuario"
)

//const Porta = ":8080"

func main() {
	db := usuario.DaoCriaBanco()

	helper.CriarDiretorioSeNaoExistir("config")

	configuracoes := config.AbrirConfiguracoes()
	porta := fmt.Sprintf(":%s", configuracoes["porta"])
	host := configuracoes["host"]

	router := route.NewRouter()

	fmt.Printf("Servidor: http://%s%s\n\n", host, porta)

	log.Fatal(http.ListenAndServe(porta, router))

	db.Close()
}

// func main() {
// fmt.Println("Teste")

// for _, usuario := range usuario.ListaUsuarios {
// 	fmt.Println(usuario)
// }

// teste1, err := usuario.DaoProcuraUsuario(2)
// fmt.Println("Usuário id 2:", teste1, err)
// teste2, err := usuario.DaoProcuraUsuario(5)
// fmt.Println("Usuário id 5:", teste2, err)

// usuario.DaoRemoveUsuario(3)
// fmt.Println("Usuário 3 removido")

// usuario.DaoRemoveUsuario(5)
// fmt.Println("Usuário 5:", err)

// for _, usuario := range usuario.ListaUsuarios {
// 	fmt.Println(usuario)
// }

// err = usuario.DaoAlteraUsuario(usuario.Usuario{Id: 2, Nome: "Francisco"})
// fmt.Println("Usuário 2 alterado")

// err = usuario.DaoAlteraUsuario(usuario.Usuario{Id: 6, Nome: "João"})
// fmt.Println("Usuário 6:", err)

// for _, usuario := range usuario.ListaUsuarios {
// 	fmt.Println(usuario)
// }

// var i int
// var p *int
// i = 42
// p = &i
// fmt.Println(i, p, *p)
// }
