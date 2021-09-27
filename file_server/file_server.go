// Arquivo file_server.go é um exemplo de servidor de arquivos em protocolo http
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// PORT define a porta padrão, atualmente 9000
const PORT = 9000

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func printLogError(msg error) {
	log.Fatal("ERROR: ", msg)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		printLogError(err)
	}

	if len(os.Args) > 1 {

		path := os.Args[1]

		ok, err := isDirectory(path)
		if err != nil {
			printLogError(err)
		}

		if ok {
			dir = path
		}
	}

	fmt.Println(dir)

	http.Handle("/", http.FileServer(http.Dir(dir)))

	formated_port := fmt.Sprintf(":%d", PORT)
	err = http.ListenAndServe(formated_port, nil)
	log.Fatal(err)
}
