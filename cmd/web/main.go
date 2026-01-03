package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// command line flag for http network address
	// default value of :4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// create logger for info messages
	// has extra flags for any more info to add
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// create logger for error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		errorLog.Fatal(err)
	}
}
