package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

func main() {
	var n int64 = 3
	var err error
	fat := fatorial // ou // var fat func(int64) *big.Int = fatorial

	if len(os.Args) == 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			help(n)
			return
		} else if n, err = converteStringToInt64(os.Args[1]); err != nil {
			panic(err.Error() + "\nParâmetro deve ser um inteiro")
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] != "-r" && os.Args[1] != "-s" {
			panic(fmt.Errorf("Parâmetro %s inválido, -h para ajuda", os.Args[1]))
		} else if n, err = converteStringToInt64(os.Args[2]); err != nil {
			panic(err.Error() + "\nSegundo parâmetro deve ser um inteiro")
		}

		fat = fatorialR
	}

	if n < 0 {
		if os.Args[1] == "-s" {
			fmt.Printf("-%d\n", fat(n))
		} else {
			fmt.Printf("-(%d!): -%d\n", -1*n, fat(n))
		}
	} else {
		if os.Args[1] == "-s" {
			fmt.Printf("%d\n", fat(n))
		} else {
			fmt.Printf("%d!: %d\n", n, fat(n))
		}
	}
}

func help(n int64) {
	fmt.Printf(`Sintaxe: %s [ OPÇÃO ] [ inteiro ]
Opções:
				Retorna o fatorial de %d
	inteiro			Retorna o fatorial no formato inteiro!: fatorial
	-r inteiro		Retorna o fatorial no formato anterior, mas usando função recursiva
	-s inteiro      Retorna o fatorial em modo simples, somente um número inteiro
	-h,  --help		Mostra esta ajuda
`, os.Args[0], n)
}

func abs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

func converteStringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64) // converte string em um int64
}

func fatorial(num int64) *big.Int {
	num = abs(num)
	fat := big.NewInt(1)
	for ; num > 1; num-- {
		fat.Mul(fat, big.NewInt(num))
	}
	return fat
}

func fatorialR(num int64) *big.Int {
	num = abs(num)
	if num == 1 {
		return big.NewInt(1)
	}

	n := big.NewInt(num)
	return n.Mul(n, fatorialR(num-1))
}

func fatorialR2(num *big.Int) *big.Int {
	num.Abs(num)
	zero := big.NewInt(0)
	um := big.NewInt(1)
	// if num.Cmp(zero) == -1 || num.Cmp(um) == 0 {
	if num.Cmp(zero) == 0 {
		return big.NewInt(1)
	}

	return num.Mul(num, fatorialR2(num.Sub(num, um)))
}
