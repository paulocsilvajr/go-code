// Package list implementa uma lista nos moldes de list em python usando slice.
package list

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// elemento é o tipo base para List.
type element interface{}

// SliceElements é uma abstração para slice de element
type SliceElements []element

// List é uma struct contendo um slice de elementos.
// Segue os moldes da implementação em Python, possuindo métodos e retornos equivalentes.
// Pode guardar elementos de qualquer tipo.
type List struct {
	Elements SliceElements
}

// Int converte um element em um int, se possível
func Int(e element) int {
	return e.(int)
}

// String converte um element em um string, se possível
func String(e element) string {
	return e.(string)
}

// Float converte um element em um float64, se possível
func Float(e element) float64 {
	return e.(float64)
}

// MakeList retorna um nova List.
// l := list.MakeList()
func MakeList() List {
	return List{}
}

// String sobrescreve interface String padrão para formatar impressão de List para:
// [0, 1, 2, ...]
func (list List) String() string {
	temp := make([]string, list.Length())
	for i, e := range list.Elements {
		temp[i] = fmt.Sprintf("%v", e)
	}

	return fmt.Sprintf("[%v]", strings.Join(temp, ", "))
}

// Length retorna o comprimento da List.
func (list *List) Length() int {
	return len(list.Elements)
}

func (list *List) Capacity() int {
	return cap(list.Elements)
}

// last retorna o índice do último elemento de list
func (list *List) last() int {
	return list.Length() - 1
}

// posicaoIndiceNegativo retorna o índice referente ao indice negativo informado
func (list *List) posicaoIndiceNegativo(index int) int {
	if index < 0 {
		return list.Length() + index
	}
	return 0
}

// Get retorna o elemento referente ao indice informado e true caso exista elemento no índice.
// Pode-se informar números negativos para elementos em ordem inversa.
// Ex: -1 == último.
func (list *List) Get(index int) (element, bool) {
	if list.last() >= 0 && index < list.Length() {
		i := index
		if index < 0 {
			i = list.posicaoIndiceNegativo(index)
		}

		return list.Elements[i], true
	}

	return 0, false
}

// GetI retorna o elemento referente ao indice informado convertido em int e true caso exista elemento no índice.
func (list *List) GetI(index int) (int, bool) {
	if e, ok := list.Get(index); ok {
		return Int(e), ok
	}
	return 0, false
}

// GetF retorna o elemento referente ao indice informado convertido em float e true caso exista elemento no índice.
func (list *List) GetF(index int) (float64, bool) {
	if e, ok := list.Get(index); ok {
		return Float(e), ok
	}
	return 0.0, false
}

// GetS retorna o elemento referente ao indice informado convertido em string e true caso exista elemento no índice.
func (list *List) GetS(index int) (string, bool) {
	if e, ok := list.Get(index); ok {
		return String(e), ok
	}
	return "", false
}

// Gets retorna o intervalo entre index0 e indexN como um slice de elementos e true caso o intervalo exista. Pode-se usar números negativos para indexN.
func (list *List) Gets(index0, indexN int) (SliceElements, bool) {
	if indexN < 0 {
		i := list.posicaoIndiceNegativo(indexN)
		return list.Elements[index0:i], true
	} else if index0 < indexN && index0 >= 0 && indexN <= list.Length() {
		return list.Elements[index0:indexN], true
	}

	return SliceElements{}, false
}

// Set atribui no índice o elemento informado.
func (list *List) Set(index int, e element) bool {
	if index < 0 {
		i := list.posicaoIndiceNegativo(index)
		list.Elements[i] = e

		return true

	} else if index < list.Length() {
		list.Elements[index] = e

		return true
	}

	return false
}

// Append adiciona o elemento informado no final de List.
func (list *List) Append(e element) {
	list.Elements = append(list.Elements, e)
}

// Appends adiciona os elementos informados no final de List.
func (list *List) Appends(e ...element) {
	list.Elements = append(list.Elements, e...)
}

// GetType retorna o tipo do elemento no índice informado.
func (list *List) GetType(index int) string {
	e, ok := list.Get(index)

	if ok {
		return fmt.Sprintf("%T", e)
	}

	return ""
}

// Extend adiciona o slice informado no final de List.
func (list *List) Extend(slice SliceElements) {
	list.Elements = append(list.Elements, slice...)
}

// Clear limpa a List.
func (list *List) Clear() {
	list.Elements = nil
}

// Copy retorna uma cópia de List.
func (list *List) Copy() List {
	cp := MakeList()
	cp.Elements = list.Elements

	return cp
}

// Count retorna a quantidade do elemento informado.
func (list *List) Count(e element) int {
	n := 0
	for _, elem := range list.Elements {
		if elem == e {
			n++
		}
	}

	return n
}

// Index retorna o indice do elemento informado e true caso o elemento exista.
func (list *List) Index(e element) (int, bool) {
	for index, elem := range list.Elements {
		if elem == e {
			return index, true
		}
	}

	return 0, false
}

// Insert incrementa(realoca elementos) em List um elemento no índice informado.
// Retorna true caso a operação ocorra com sucesso.
func (list *List) Insert(index int, e element) bool {
	if index == 0 || index < list.Length() {
		var temp SliceElements
		temp = append(temp, list.Elements[index:]...)
		list.Elements = append(list.Elements[:index], e)
		list.Elements = append(list.Elements, temp...)

		return true
	} else {
		return false
	}

}

// Pop remove o elemento de List no índice informado.
// Retorna o elemento removido e true caso operação ocorra com sucesso.
func (list *List) Pop(index ...int) (element, bool) {
	t := list.Length()
	var elem element
	last := list.last()

	if len(index) == 0 {
		if t == 0 {
			return elem, false
		}

		elem = list.Elements[last]
		list.Elements = list.Elements[:last]

		return elem, true

	} else if index[0] >= 0 && index[0] <= last {
		i0 := index[0]
		elem = list.Elements[i0]
		list.Elements = append(list.Elements[:i0], list.Elements[i0+1:]...)
		return elem, true
	}

	return elem, false
}

// Remove remove o elemento informado de List.
// Retorna true caso operação ocorra com sucesso.
func (list *List) Remove(e element) bool {
	for i := 0; i < list.Length(); i++ {

		if list.Elements[i] == e {
			_, ok := list.Pop(i)
			return ok
		}
	}

	return false
}

// Sort ordena List em ordem crescente. Obrigatório informar sortFunction para tipos não primitivos.
//
// Implemente os sequintes metodos/funções/tipos caso o conteúdo de List(list.Elements)
// não seja dos tipos primitivos(int, float64, string):
//
// type tipoOrdenacao list.SliceElements
//
// func (n tipoOrdenacao) Len() int           { return len(n) }
//
// func (n tipoOrdenacao) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
//
// func (n tipoOrdenacao) Less(i, j int) bool { return n[i].(TipoConteudo).atributo < n[j].(TipoConteudo).atributo }
//
// func (p TipoConteudo) String() string   { return fmt.Sprintf("%v %v", p.atributo, p.atributo2) } // pode-se omitir, caso não queira personalizar a saída de impressão
//
// func sortFunction(list *list.List) {
// 	sort.Sort(tipoOrdenacao(list.Elements))
// }
func (list *List) Sort(sortFunction ...func(*List)) error {
	tp := make(map[string]bool)

	for i := range list.Elements {
		tp[list.GetType(i)] = true
	}

	if len(tp) > 1 {
		return errors.New("Lista heterogênea não pode ser ordenada")

	} else {

		if list.Length() > 0 {
			switch list.Elements[0].(type) {
			case int:
				var ti []int

				for i := range list.Elements {
					ti = append(ti, list.Elements[i].(int))
				}

				sort.Ints(ti)

				for i, e := range ti {
					list.Elements[i] = e
				}

			case float64:
				var tf []float64

				for i := range list.Elements {
					tf = append(tf, list.Elements[i].(float64))
				}

				sort.Float64s(tf)

				for i, e := range tf {
					list.Elements[i] = e
				}

			case string:
				var ts []string

				for i := range list.Elements {
					ts = append(ts, list.Elements[i].(string))
				}

				sort.Strings(ts)

				for i, e := range ts {
					list.Elements[i] = e
				}

			default:
				if len(sortFunction) == 1 {
					function := sortFunction[0]
					function(list)

				} else {
					return errors.New("Parâmetro sortFunction inválido")
				}
			}
		}

		return nil
	}
}

// Reverse ordena List em ordem decrescente. Mais detalhes em godoc de Sort().
func (list *List) Reverse(sortFunction ...func(*List)) error {
	var er error
	if len(sortFunction) == 1 {
		function := sortFunction[0]
		er = list.Sort(function)
	} else {
		er = list.Sort()
	}

	for i, j := 0, list.Length()-1; i < list.Length()/2; i, j = i+1, j-1 {
		list.Elements[i], list.Elements[j] = list.Elements[j], list.Elements[i]
	}

	return er
}
