package models

import (
	"database/sql"
	"time"
)

// Define a CLip type to hold the data for an individual clip
type Clip struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define clipModel type which wraps a sql.DB connection pool
type ClipModel struct {
	DB *sql.DB
}

// This will insert a new clip into the database
func (m *ClipModel) Insert(title string, content string, expires int) (int, error) {
	var id int
	stmt := `INSERT INTO vaults (title, content, created, expires)
	VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 day' * $3) RETURNING id`

	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

// This will return a specific clips based on its id
func (m *ClipModel) Get(id int) (*Clip, error) {
	return nil, nil
}

// This will return the 10 most recent created clips
func (m *ClipModel) Latest() ([]*Clip, error) {
	return nil, nil
}
