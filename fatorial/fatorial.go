package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	var n int64 = 3

	if len(os.Args) > 1 {
		var err error
		n, err = strconv.ParseInt(os.Args[1], 10, 64) // converte string em um int64
		if err != nil {
			panic(err.Error() + "\nParâmetro deve ser um inteiro")
		}
	}

	fmt.Printf("%d!: %d\n", n, fat(n))
}

func fat(num int64) *big.Int {
	fat := big.NewInt(1)
	for ; num > 1; num-- {
		fat.Mul(fat, big.NewInt(num))
	}
	return fat
}
