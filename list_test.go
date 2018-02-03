package main

import (
	"fmt"
	"go-code/list"
	"reflect"
	"sort"
	"testing"
)

var l1 list.List
var l2 list.List

func TestMakeList(t *testing.T) {
	l := list.MakeList()

	tipo := fmt.Sprintf("%v", reflect.TypeOf(l))
	if tipo != "list.List" || !reflect.DeepEqual(l, list.List{}) {
		t.Errorf("MakeList não está retornando uma list.List, retorna uma %v",
			tipo)
	}

	fmt.Println("list.MakeList")
}

func TestString(t *testing.T) {
	if fmt.Sprint(l1) != "[]" {
		t.Errorf("String não está retornando o formato padrão estabelecido quando vazio.")
	}

	l1.Appends(1)
	if fmt.Sprint(l1) != "[1]" {
		t.Errorf("String não está retornando o formato padrão estabelecido quando um elemento.")
	}

	l1.Appends(2, 3)
	if fmt.Sprint(l1) != "[1, 2, 3]" {
		t.Errorf("String não está retornando o formato padrão estabelecido quando possui vários elementos.")
	}

	fmt.Println("list.String")
}

func TestLength(t *testing.T) {
	q := 3
	if l1.Length() != q {
		t.Errorf("Length não está retornando o valor correto para a quantidade de elementos. %d != %d", l1.Length(), q)
	}

	fmt.Println("list.Length")
}

func TestElements(t *testing.T) {
	e := l1.Elements[1]
	n := 2
	if e != n {
		t.Errorf("Elements[1] não está retornando o valor correto, %d != %d.", e, n)
	}

	es := l1.Elements[1:3]
	ns := list.SliceElements{2, 3}
	if !reflect.DeepEqual(es, ns) {
		t.Errorf("Elements[1:2] não está retornando o valor correto, %v[%s] != %v[%s].", es, reflect.TypeOf(es), ns, reflect.TypeOf(ns))
	}

	fmt.Println("list.Elements")
}

func TestGetGets(t *testing.T) {
	e, ok := l1.Get(1)
	n := 2
	if e != n {
		t.Errorf("Get(1) não está retornando o valor correto, %d[%t] != %d.", e, ok, n)
	}

	es, oks := l1.Gets(1, 3)
	ns := list.SliceElements{2, 3}
	if !reflect.DeepEqual(es, ns) {
		t.Errorf("Gets(1, 2) não está retornando o valor correto, %v[%t][%s] != %v[%s].",
			es, oks, reflect.TypeOf(es), ns, reflect.TypeOf(ns))
	}

	e, ok = l1.Get(-1)
	n = 3
	if e != n {
		t.Errorf("Get(-1) não está retornando o valor correto, %d[%t] != %d.", e, ok, n)
	}

	if e, ok = l1.Get(3); ok {
		t.Errorf("Em e, ok := Get(3) a variável ok está retornando verdadeiro[ok=%t] para um índice inválido", ok)
	}

	es, oks = l1.Gets(0, -1)
	ns = list.SliceElements{1, 2}
	if !reflect.DeepEqual(es, ns) {
		t.Errorf("Gets(0, -1) não está retornando o valor correto, %v[%t][%s] != %v[%s].",
			es, oks, reflect.TypeOf(es), ns, reflect.TypeOf(ns))
	}

	var i int
	if i, ok = l1.GetI(3); fmt.Sprintf("%s", reflect.TypeOf(e)) != "int" {
		t.Errorf("Em i, ok := GetI(3) a variável i[%d] deve ser do tipo int", i)
	}

	if es, oks = l1.Gets(-1, 3); ok {
		t.Errorf("Em es, oks := Gets(-1, 3) a variável ok está retornando verdadeiro[ok=%t] para um índice inválido", ok)
	}

	if es, oks = l1.Gets(0, 4); ok {
		t.Errorf("Em es, oks := Gets(0, 4) a variável ok está retornando verdadeiro[ok=%t] para um índice inválido", ok)
	}

	fmt.Println("list.Get/Gets")
}

func TestSet(t *testing.T) {
	if ok := l1.Set(2, 5); !ok {
		t.Errorf("Set(2, 5) está retornando falso[ok=%t] para um índice válido", ok)
	}

	if ok := l1.Set(3, 5); ok {
		t.Errorf("Set(3, 5) está retornando verdadeiro[ok=%t] para um índice inválido", ok)
	}

	if ok := l1.Set(-1, 3); !ok {
		t.Errorf("Set(-1, 3) está retornando falso[ok=%t] para um índice válido(inverso)", ok)
	}

	fmt.Println("list.Set")
}

func TestAppendAppends(t *testing.T) {
	l1.Append(4)
	l := list.List{list.SliceElements{1, 2, 3, 4}}
	if !reflect.DeepEqual(l1, l) {
		t.Errorf("Append(4) está adicionando elementos errado")
	}

	l1.Appends(5, 6)
	l.Elements = append(l.Elements, 5, 6)
	if !reflect.DeepEqual(l1, l) {
		t.Errorf("Appends(5, 6) está adicionando elementos errado")
	}

	l1.Appends(list.SliceElements{7, 8}...)
	l.Elements = append(l.Elements, list.SliceElements{7, 8}...)
	if !reflect.DeepEqual(l1, l) {
		t.Errorf("l1.Appends(list.SliceElements{7, 8}...) está adicionando elementos errado")
	}

	fmt.Println("list.Append/Appends")
}

func TestGetType(t *testing.T) {
	if l1.GetType(0) != "int" {
		t.Errorf("GetType(0) está retornando uma string com o tipo errado")
	}

	if l1.GetType(-2) != "int" {
		t.Errorf("GetType(-2) está retornando uma string com o tipo errado")
	}

	if l1.GetType(10) != "" {
		t.Errorf("GetType(10) está retornando uma string com um tipo para uma posição inválida")
	}

	fmt.Println("list.GetType")
}

func TestExtend(t *testing.T) {
	l1.Extend(list.SliceElements{9, 10})
	l := list.List{list.SliceElements{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}

	if !reflect.DeepEqual(l1, l) {
		t.Errorf("Extend(list.SliceElements{9, 10}) está adicionando SliceElements errado")
	}

	fmt.Println("list.Extend")
}

func TestCopy(t *testing.T) {
	l2 = l1.Copy()

	if !reflect.DeepEqual(l1, l2) {
		t.Errorf("Copy() está retornado uma cópia inválida, %v != %v", l1, l2)
	}

	fmt.Println("list.Copy")
}

func TestClear(t *testing.T) {
	l2.Clear()

	if !(l2.Length() > 0) && !reflect.DeepEqual(l2, list.List{}) {
		t.Errorf("Clear() está limpando a List incorretamente")
	}

	fmt.Println("list.Clear")
}

func TestCount(t *testing.T) {
	l1.Set(-1, 9)
	if l1.Count(9) != 2 && l1.Count(2) != 1 {
		t.Errorf("Count(9)/Count(1) está retornando a quantidade errada do elemento informado")
	}

	if l1.Count(10) > 0 {
		t.Errorf("Count(10) está retornando quantidade para elemento inexistente em List")
	}

	fmt.Println("list.Count")
}

func TestIndex(t *testing.T) {
	i, ok := l1.Index(2)
	if i != 1 && !ok {
		t.Errorf("Index(2) está retornando uma posição inválida")
	}

	if i, ok = l1.Index(10); ok {
		t.Errorf("Index(10) está retornando uma posição para um elemento inexistente")
	}

	fmt.Println("list.Index")
}

func TestInsert(t *testing.T) {
	l := list.List{list.SliceElements{10}}
	ok := l2.Insert(0, 10)
	if !ok && reflect.DeepEqual(l2, l) {
		t.Errorf("Insert(0, 10) não esta inserindo no indice 0 o elemento 10")
	}

	if ok = l2.Insert(5, 20); ok {
		t.Errorf("Insert(5, 20) está inserindo o elemento 20 em uma posição inexistente")
	}

	fmt.Println("list.Insert")
}

func TestPop(t *testing.T) {
	l2.Insert(0, 20)

	o := l2.Elements[1]
	e, ok := l2.Pop(1)
	if !ok && e != o {
		t.Errorf("Pop(1) não está removendo elemento da posição 1")
	}

	o = l2.Elements[l2.Length()-1]
	e, ok = l2.Pop()
	if !ok && e != o {
		t.Errorf("Pop() não está removendo o último elemento da List")
	}

	e, ok = l2.Pop()
	if ok {
		t.Errorf("Pop() está removendo elementos de List vazia")
	}

	fmt.Println("list.Pop")
}

func TestRemove(t *testing.T) {
	if ok := l1.Remove(9); !ok {
		t.Errorf("Remove(9) não está removendo elemento existente na List")
	}

	if ok := l1.Remove(10); ok {
		t.Errorf("Remove(10) está removendo elemento inexistente da List")
	}

	fmt.Println("list.Remove")
}

// tipo, struct e métodos para ordenação de Pessoa
type Pessoa struct {
	nome  string
	idade int
}
type sorterIdade list.SliceElements

func (n sorterIdade) Len() int           { return len(n) }
func (n sorterIdade) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n sorterIdade) Less(i, j int) bool { return n[i].(Pessoa).idade < n[j].(Pessoa).idade }

func sortFunction(list *list.List) {
	sort.Sort(sorterIdade(list.Elements))
}

func TestSortReverse(t *testing.T) {
	er := l1.Sort()
	if er != nil {
		t.Errorf("Sort() não pode ordenar a List de inteiros")
	}

	var l3 list.List
	l3.Appends(list.SliceElements{"c", "a", "b", "f", "e", "d"}...)
	er = l3.Sort()
	if er != nil {
		t.Errorf("Sort() não pode ordenar a List de strings")
	}

	var l4 list.List
	l4.Appends(list.SliceElements{
		Pessoa{"Paulo", 30},
		Pessoa{"João", 20},
		Pessoa{"André", 25}})
	er = l4.Sort(sortFunction)
	if er != nil {
		t.Errorf("Sort() não pode ordenar a List de strings")
	}

	var l5 list.List
	l5.Appends(list.SliceElements{"c", 2, 3.14}...)
	er = l5.Sort()
	if er == nil {
		t.Errorf("Sort() em List heterogênea deve retornar erro")
	}

	er = l1.Reverse()
	if er != nil {
		t.Errorf("Reverse() não pode ordenar em ordem inversa a List de inteiros")
	}

	fmt.Println("list.Sort/Reverse")
}
