package main

import (
	"fmt"

	"github.com/paulocsilvajr/go-code/list"
)

func main() {
	fmt.Println("Exemplo de pacote list")

	l := list.MakeList()
	l.Extend(list.SliceElements{1, 1, 2, 3, 5, 8, 13})
	ln, _ := l.GetI(-1)
	ln1, _ := l.GetI(-2)
	proximo := ln + ln1
	l.Append(proximo)

	fmt.Println(l)

	fmt.Println(len(l.Elements), cap(l.Elements))
}
