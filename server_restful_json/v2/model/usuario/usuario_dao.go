package usuario

import (
	"errors"
	"fmt"
	"time"
)

var idCorrente int

var ListaUsuarios Usuarios

func init() {
	DaoAdicionaUsuario(
		Usuario{
			Nome:        "paulo",
			Email:       "pauluscave@gmail.com",
			Ativo:       true,
			DataCriacao: time.Now()})

	DaoAdicionaUsuario(
		Usuario{
			Nome:        "jose",
			Email:       "jose@gmail.com",
			Senha:       "123456",
			Ativo:       true,
			DataCriacao: time.Now()})

	DaoAdicionaUsuario(
		Usuario{
			Nome:  "andré",
			Email: "andre@gmail.com",
			Ativo: true})
}

func DaoAdicionaUsuario(usuario Usuario) Usuario {
	idCorrente++
	usuario.Id = idCorrente
	ListaUsuarios = append(ListaUsuarios, &usuario)

	return usuario
}

func DaoRemoveUsuario(id int) error {
	for i, usuario := range ListaUsuarios {
		if usuario.Id == id {
			ListaUsuarios = append(ListaUsuarios[:i], ListaUsuarios[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Não foi encontrado usuário com id[%d]", id)
}

func DaoAlteraUsuario(usuario Usuario) Usuario {
	for _, usuarioAtual := range ListaUsuarios {
		fmt.Println(usuarioAtual)
		if usuario.Id == usuarioAtual.Id {
			usuarioAtual.Nome = usuario.Nome
			usuarioAtual.Email = usuario.Email
			usuarioAtual.Senha = usuario.Senha
			usuarioAtual.Ativo = usuario.Ativo
			usuarioAtual.DataCriacao = usuario.DataCriacao

			return *usuarioAtual
		}
	}

	return Usuario{}
}

func DaoProcuraUsuario(id int) (*Usuario, error) {
	for _, usuario := range ListaUsuarios {
		if usuario.Id == id {
			return usuario, nil
		}
	}

	return &Usuario{}, errors.New("Usuário não existe")
}
