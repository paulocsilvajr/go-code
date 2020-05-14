package util

import "testing"

type Pessoa struct {
	Nome  string
	idade int
}

func (p Pessoa) Andar() string {
	return "Andando..."
}

func (p Pessoa) Correr() string {
	return "Correndo..."
}

func (p Pessoa) levitar() string {
	return "Levitando..."
}

func TestType(t *testing.T) {
	inteiro := 42
	tipo := Type(inteiro)
	if tipo != "int" {
		t.Errorf("Type retornando tipo incorreto '%s'", tipo)
	}

	p := Pessoa{"João", 25}
	tipo = Type(p)
	if tipo != "Pessoa" {
		t.Errorf("Type retornando tipo incorreto '%s'", tipo)
	}
}

func TestDir(t *testing.T) {
	p := Pessoa{"José", 30}
	metodos := Dir(p)
	if len(metodos) > 2 {
		t.Error("Dir retornou um slice com mais de 2 elemento. Struct Pessoa tem somente 2 método Públicos", metodos)
	}

	if metodos[0] != "Andar" {
		t.Error("Dir não retornou o método declarado na struct Pessoa chamado Andar")
	}

	if metodos[1] != "Correr" {
		t.Error("Dir não retornou o método declarado na struct Pessoa chamado Correr")
	}
}
