package models

import (
	"database/sql"
	"time"
)

// Snippet holds data for individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert adds a new snippet to the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get returns a specific snippet based on id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest returns the 10 recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
