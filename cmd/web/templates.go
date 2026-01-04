package main

import "github.com/AutomationMK/snippetbox/internal/models"

// define template data struct to hold dynamic data
// for html templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
