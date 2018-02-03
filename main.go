package main

import (
	"fmt"
	"go-code/list"
)

func main() {
	fmt.Println("Testes de pacotes")

	l := list.MakeList()
	l.Extend(list.SliceElements{1, 1, 2, 3, 5, 8, 13})
	ln, _ := l.Get(-1)
	ln_1, _ := l.Get(-2)
	proximo := ln.(int) + ln_1.(int)
	l.Append(proximo)

	fmt.Println(proximo)

	n1 := l.Elements[0]
	n2, _ := l.GetI(-2)

	n1 = n1.(int) + n2

}
