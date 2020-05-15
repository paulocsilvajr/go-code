package util

import (
	"reflect"
)

// DirM retorna um slice de strings contendo os métodos Públicos da struct informada. Retorna um slice vazio se a variável informada não for uma struct
func DirM(v interface{}) []string {
	tipo := getType(v)

	var metodos []string
	for i := 0; i < tipo.NumMethod(); i++ {
		m := tipo.Method(i)
		metodos = append(metodos, m.Name)
	}

	return metodos
}

// DirF retorna um slice de strings contendo os atributos da struct informada. Retorna um slice vazio se a variável informada não for uma struct
func DirF(v interface{}) []string {
	tipo := getType(v)

	if tipo.Kind() != reflect.Struct {
		return []string{}
	}

	var atributos []string
	for i := 0; i < tipo.NumField(); i++ {
		a := tipo.Field(i)
		atributos = append(atributos, a.Name)
	}

	return atributos
}

// Dir retorna um map contendo os métodos(chave: methods) e atributos(chave: fields) da struct informada.
func Dir(v interface{}) map[string][]string {
	metodos := DirM(v)
	atributos := DirF(v)

	dir := make(map[string][]string)
	dir["methods"] = metodos
	dir["fields"] = atributos

	return dir
}

// Type retorna uma string contendo o tipo da variável informada
func Type(v interface{}) string {
	t := getType(v)

	return t.Name()
}

// getType retorna o tipo(reflect.Type) da variável passada como parâmetro
func getType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
