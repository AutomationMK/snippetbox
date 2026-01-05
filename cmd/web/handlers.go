package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AutomationMK/snippetbox/internal/models"
)

// home is the home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if route is exactly the home route
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// render the home template page
	app.render(w, http.StatusOK, "home.tmpl", &templateData{
		Snippets: snippets,
	})
}

// snippetView is the snippet view handler function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// extract value from query string
	// convert that value to an integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// use the SnippetModel get method to retrive
	// the data for a specific record based on
	// it's id. if no record then return 404
	// not found http response
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// parse the view template file
	app.render(w, http.StatusOK, "view.tmpl", &templateData{
		Snippet: snippet,
	})
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
