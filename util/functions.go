package util

import (
	"reflect"
)

// Dir retorna um slice de strings contendo os métodos Públicos da classe informada
func Dir(v interface{}) []string {
	tipo := reflect.TypeOf(v)

	var metodos []string
	for i := 0; i < tipo.NumMethod(); i++ {
		m := tipo.Method(i)
		metodos = append(metodos, m.Name)
	}

	return metodos
}

// Type retorna uma string contendo o tipo da variável informada
func Type(v interface{}) string {
	t := reflect.TypeOf(v)

	return t.Name()
}
