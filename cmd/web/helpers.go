package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// render helps render a template page
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// retrieve the appropriate template set based
	// on page name
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	// initialize a new buffer
	buf := new(bytes.Buffer)

	// write the template to the buffer and check
	// for any errors and if so do not serve html
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// write out the provided HTTP status code
	w.WriteHeader(status)

	// write the contents of the buffer to the http.ResponseWriter
	buf.WriteTo(w)
}

// serverError writes an error message and the
// stack trace to the errorLog. then sends a
// generic 500 internal server error response
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code
// and corresponding description to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound uses the clientError function to
// send a notFound 404 response
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
