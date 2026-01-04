package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home is the home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if route is exactly the home route
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// initialize a slice for the two templates
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	// parse the home page template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// write the template to the response
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// snippetView is the snippet view handler function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// extract value from query string
	// convert that valute to and integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate is the snippet create handler function
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check if request is a POST request
	if r.Method != http.MethodPost {
		// send a 405 status along with message
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// create some dummy variables for testing
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// insert the dummy data into the database
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect the user to the relevant snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
