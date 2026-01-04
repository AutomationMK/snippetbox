package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/AutomationMK/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// struct to hold application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// command line flag for http network address
	// default value of :4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// command line flag for database connection
	dsn := flag.String("dsn", "snippetbox:pass@/snippetbox?parseTime=true", "MySQL data source name")

	// create logger for info messages
	// has extra flags for any more info to add
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// create logger for error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// use openDB() function passing the DSN
	// from the command line flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// application struct containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// initialize a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

// openDB() waps sql.Open() and returns a sql.DB
// connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	// initialize the pool for future use
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// create a simple connection to check for errors
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
