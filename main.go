package main

import (
	"fmt"
	"go-code/list"
)

func main() {
	fmt.Println("Testes de pacotes")

	l := list.MakeList()
	l.Extend(list.SliceElements{1, 1, 2, 3, 5, 8, 13})
	ln, _ := l.GetI(-1)
	ln_1, _ := l.GetI(-2)
	proximo := ln + ln_1
	l.Append(proximo)

	fmt.Println(l)

	fmt.Println(len(l.Elements), cap(l.Elements))

}
