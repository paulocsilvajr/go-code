package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// PORT porta que será disponibilizada pelo servidor
const PORT = 3000

func main() {
	if len(os.Args) > 1 {
		path := fmt.Sprintf("./%s", os.Args[1])

		if stat, err := os.Stat(path); err == nil && stat.IsDir() {
			http.Handle("/", http.FileServer(http.Dir(path)))
			url := fmt.Sprintf(":%d", PORT)
			fmt.Printf("Rodando servidor em localhost:%d\nCTRL+C para finalizar\n", PORT)

			err := http.ListenAndServe(url, nil)

			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("Caminho '%s' não é um diretório válido\n", path)
		}
	} else {
		fmt.Println("Informe como parâmetro um diretório para o servidor.")
	}
}
