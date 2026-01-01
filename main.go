package main

import (
	"log"
	"net/http"
)

// home is the home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet Box"))
}

// snippetView is the snippet view handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet"))

}

// snippetCreate is the snippet create handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	port := ":4000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetCreate)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
