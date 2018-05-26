package usuario

import (
	"time"
)

type Usuario struct {
	Id          int       `json:"id"`
	Nome        string    `json:"nome"`
	Email       string    `json:"email"`
	Senha       string    `json:"senha"`
	Ativo       bool      `json:"ativo"`
	DataCriacao time.Time `json:"data_criacao"`
}

type Usuarios []*Usuario
