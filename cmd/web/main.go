package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/codebox", codebox)
	mux.HandleFunc("/codebox/create", createCodebox)

	log.Println("Server hosted on localhost: 4001")
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		panic(err)
	}
}
