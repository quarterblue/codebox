package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/codebox", app.codebox)
	mux.HandleFunc("/codebox/create", app.createCodebox)

	fileServer := http.FileServer(http.Dir(".ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}