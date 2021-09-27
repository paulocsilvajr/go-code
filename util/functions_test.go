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

func (p Pessoa) GetIdade() int {
	return p.idade
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
	metodos := DirM(p)
	if len(metodos) > 3 {
		t.Error("DirM retornou um slice com mais de 3 elemento. Struct Pessoa tem somente 3 método Públicos", metodos)
	}

	for i := 0; i < len(metodos); i++ {
		metodo := metodos[i]
		switch metodo {
		case "Andar", "Correr", "GetIdade":
			continue
		default:
			t.Errorf("DirM não retornou o método declarado na struct Pessoa chamado %s", metodo)
		}
	}

	atributos := DirF(p)
	if len(atributos) > 2 {
		t.Error("DirF retornou um slice com mais de 2 elementos. Struct Pessoa tem somente 2 atributos", atributos)
	}

	if atributos[0] != "Nome" {
		t.Error("DirF não retornou o atributo declarado na struct Pessoa chamado Nome")
	}

	if atributos[1] != "idade" {
		t.Error("DirF não retornou o atributo declarado na struct Pessoa chamado idade")
	}

	metodos = DirM(42)
	if len(metodos) > 0 {
		t.Error("DirM retornou métodos para uma variável inteira", metodos)
	}

	atributos = DirF(42)
	if len(atributos) > 0 {
		t.Error("DirF retornou atributos para uma variável inteira", atributos)
	}

	dir := Dir(p)
	if len(dir["methods"]) > 2 && len(dir["fields"]) > 2 {
		t.Error("Dir retornou mais de 2 métodos e 2 atributos. Struct Pessoa tem 2 métodos publicos e 2 atributos", dir)
	}
}
