package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":4000"
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
