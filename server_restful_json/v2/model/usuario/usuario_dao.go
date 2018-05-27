package usuario

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var ListaUsuarios Usuarios
var db *sql.DB

func DaoCriaBanco() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./banco.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	sql := `
		CREATE TABLE IF NOT EXISTS usuario(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			nome TEXT NOT NULL,
			email TEXT NOT NULL,
			senha TEXT NOT NULL,
			ativo INTEGER NOT NULL,
			data_criacao TIMESTAMP NOT NULL);
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Print(err)
	}

	// DaoAdicionaUsuario(
	// 	Usuario{
	// 		Nome:        "paulo",
	// 		Email:       "pauluscave@gmail.com",
	// 		Ativo:       true,
	// 		DataCriacao: time.Now()})

	// DaoAdicionaUsuario(
	// 	Usuario{
	// 		Nome:        "jose",
	// 		Email:       "jose@gmail.com",
	// 		Senha:       "123456",
	// 		Ativo:       true,
	// 		DataCriacao: time.Now()})

	// DaoAdicionaUsuario(
	// 	Usuario{
	// 		Nome:  "andré",
	// 		Email: "andre@gmail.com",
	// 		Ativo: true})

	return db
}

func DaoCarregaUsuarios() {
	ListaUsuarios = nil
	rows, err := db.Query("SELECT id, nome, email, senha, ativo, data_criacao FROM usuario")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		usuarioAtual := Usuario{}
		err = rows.Scan(
			&usuarioAtual.Id,
			&usuarioAtual.Nome,
			&usuarioAtual.Email,
			&usuarioAtual.Senha,
			&usuarioAtual.Ativo,
			&usuarioAtual.DataCriacao)
		if err != nil {
			log.Print(err)
		}
		ListaUsuarios = append(ListaUsuarios, &usuarioAtual)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}
}

func DaoAdicionaUsuario(usuario Usuario) Usuario {
	transacao, err := db.Begin()
	if err != nil {
		log.Print(err)
	}

	stmt, err := transacao.Prepare(`
		INSERT INTO usuario(
			nome, email, senha, ativo, data_criacao) 
		VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuario.Ativo,
		usuario.DataCriacao)
	if err != nil {
		log.Print(err)
	}
	transacao.Commit()

	id, err := resultado.LastInsertId()
	if err != nil {
		log.Print(err)
	} else {
		usuario.Id = int(id)
	}

	return usuario
}

func DaoRemoveUsuario(id int) error {
	_, err := db.Exec(fmt.Sprintf("DELETE FROM usuario WHERE id = %d", id))
	if err != nil {
		log.Print(err)
	}

	return err
}

func DaoAlteraUsuario(usuario Usuario) Usuario {
	transacao, err := db.Begin()
	if err != nil {
		log.Print(err)
	}

	stmt, err := transacao.Prepare(`
		UPDATE usuario
		SET nome = ?, email = ?, senha = ?, ativo = ?, data_criacao = ? 
		WHERE id = ?`)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuario.Ativo,
		usuario.DataCriacao,
		usuario.Id)
	if err != nil {
		log.Print(err)
		usuario = Usuario{}
	}
	transacao.Commit()

	return usuario
}

func DaoProcuraUsuario(id int) (*Usuario, error) {
	rows, err := db.Query(fmt.Sprint(`
		SELECT id, nome, email, senha, ativo, data_criacao 
		FROM usuario
		WHERE id = `, id))
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var usuarioEncontrado Usuario
	for rows.Next() {
		err = rows.Scan(
			&usuarioEncontrado.Id,
			&usuarioEncontrado.Nome,
			&usuarioEncontrado.Email,
			&usuarioEncontrado.Senha,
			&usuarioEncontrado.Ativo,
			&usuarioEncontrado.DataCriacao)
		if err != nil {
			log.Print(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	return &usuarioEncontrado, nil
}

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// var idCorrente int

// var ListaUsuarios Usuarios

// func init() {
// 	DaoAdicionaUsuario(
// 		Usuario{
// 			Nome:        "paulo",
// 			Email:       "pauluscave@gmail.com",
// 			Ativo:       true,
// 			DataCriacao: time.Now()})

// 	DaoAdicionaUsuario(
// 		Usuario{
// 			Nome:        "jose",
// 			Email:       "jose@gmail.com",
// 			Senha:       "123456",
// 			Ativo:       true,
// 			DataCriacao: time.Now()})

// 	DaoAdicionaUsuario(
// 		Usuario{
// 			Nome:  "andré",
// 			Email: "andre@gmail.com",
// 			Ativo: true})
// }

// func DaoAdicionaUsuario(usuario Usuario) Usuario {
// 	idCorrente++
// 	usuario.Id = idCorrente
// 	ListaUsuarios = append(ListaUsuarios, &usuario)

// 	return usuario
// }

// func DaoRemoveUsuario(id int) error {
// 	for i, usuario := range ListaUsuarios {
// 		if usuario.Id == id {
// 			ListaUsuarios = append(ListaUsuarios[:i], ListaUsuarios[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("Não foi encontrado usuário com id[%d]", id)
// }

// func DaoAlteraUsuario(usuario Usuario) Usuario {
// 	for _, usuarioAtual := range ListaUsuarios {
// 		fmt.Println(usuarioAtual)
// 		if usuario.Id == usuarioAtual.Id {
// 			usuarioAtual.Nome = usuario.Nome
// 			usuarioAtual.Email = usuario.Email
// 			usuarioAtual.Senha = usuario.Senha
// 			usuarioAtual.Ativo = usuario.Ativo
// 			usuarioAtual.DataCriacao = usuario.DataCriacao

// 			return *usuarioAtual
// 		}
// 	}

// 	return Usuario{}
// }

// func DaoProcuraUsuario(id int) (*Usuario, error) {
// 	for _, usuario := range ListaUsuarios {
// 		if usuario.Id == id {
// 			return usuario, nil
// 		}
// 	}

// 	return &Usuario{}, errors.New("Usuário não existe")
// }
