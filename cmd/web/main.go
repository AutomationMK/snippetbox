package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// struct to hold application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	// application struct containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// initialize a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
