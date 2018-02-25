package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	router := NewRouter()

	fmt.Printf("Server: http://localhost%s\n\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, router))

}
