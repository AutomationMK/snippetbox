package main

import (
	"path/filepath"
	"text/template"

	"github.com/AutomationMK/snippetbox/internal/models"
)

// newTemplateCache creates a template cache
func newTemplateCache() (map[string]*template.Template, error) {
	// initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// get a slice of all filepaths
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// loop through the page filepaths
	for _, page := range pages {
		// extract the file name as a variable
		name := filepath.Base(page)

		// parse the base template file
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// parse all partial template files
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// parse the page template file
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add template set to the map with page name
		// as the key
		cache[name] = ts
	}

	return cache, nil
}

// define template data struct to hold dynamic data
// for html templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
