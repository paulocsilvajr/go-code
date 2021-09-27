package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// MILLISECONDS para time out
const MILLISECONDS = 1000 * time.Millisecond

func get(url string, response chan *http.Response, errors chan *error) {
	resp, err := http.Get(url)
	if err != nil {
		errors <- &err
	}
	response <- resp
}

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]

		response := make(chan *http.Response, 1)
		errors := make(chan *error)

		go get(url, response, errors)

		for {
			select {
			case r := <-response:
				fmt.Printf("%s", r.Body)
				return
			case err := <-errors:
				log.Fatal(*err)
			case <-time.After(MILLISECONDS):
				fmt.Println("Time out!")
				return
			}
		}

	} else {
		fmt.Println("Informe uma url.")
	}
}
