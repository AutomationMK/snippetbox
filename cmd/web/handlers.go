package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home is the home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// check if route is exactly the home route
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// initialize a slice for the two templates
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// parse the home page template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write the template to the response
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// snippetView is the snippet view handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract value from query string
	// convert that valute to and integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate is the snippet create handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check if request is a POST request
	if r.Method != http.MethodPost {
		// send a 405 status along with message
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
