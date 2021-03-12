package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/codebox", codebox)
	mux.HandleFunc("/codebox/create", createCodebox)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server hosted on localhost %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		panic(err)
	}
}
