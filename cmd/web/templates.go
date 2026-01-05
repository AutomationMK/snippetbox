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

		// create a slice containing the base template
		// and any patials along with the page
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		// parse the files into a template set
		ts, err := template.ParseFiles(files...)
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
