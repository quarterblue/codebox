package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	// CLI Parser to set port
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/codebox", codebox)
	mux.HandleFunc("/codebox/create", createCodebox)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Server hosted on localhost %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
