package models

import (
	"database/sql"
	"errors"
	"time"
)

type ClipModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Clip, error)
	Latest() ([]*Clip, error)
}

// Define a Clip type to hold the data for an individual clip
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
	stmt := `SELECT id, title, content, created, expires
    FROM vaults
    WHERE expires > CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &Clip{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recent created clips
func (m *ClipModel) Latest() ([]*Clip, error) {
	stmt := `SELECT id, title, content, created, expires
    FROM vaults
    WHERE expires > NOW()
    ORDER BY id DESC
    LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialize an empty slice to hold the Clip struct
	clips := []*Clip{}

	// Use rows.Next to iterate through the rows in the resultset
	for rows.Next() {
		s := &Clip{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		clips = append(clips, s)
	}

	// When the rows.Next() loop has finished, call rows.Err(0 to retrieve any error
	// that was encountered during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went Ok then return the Clips slice
	return clips, nil
}
