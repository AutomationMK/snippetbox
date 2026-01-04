package models

import (
	"database/sql"
	"errors"
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
	// write the sql insert statement
	stmt := `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// use the Exec method to execute the statement
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get the id of the new snippet
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// convert id to int type before returning
	return int(id), nil
}

// Get returns a specific snippet based on id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// write sql statement to get snippet from id
	stmt := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// use QueryRow() method to execute statement
	row := m.DB.QueryRow(stmt, id)

	// initialize a pointer to Snippet struct
	s := &Snippet{}

	// row.Scan() copies values from each field
	// in sql.Row to the snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if no rows found then row.Scan()
		// will return sql.ErrNoRows error.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// if no errors then return the snippet
	return s, nil
}

// Latest returns the 10 recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
