package models

import (
	"database/sql"
	"time"
)

// Define a Vault type to hold the data for an individual vault
type Vault struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a VaultModel type which wraps a sql.DB connection pool
type VaultModel struct {
	DB *sql.DB
}

// This will insert a new vault into the database
func (m *VaultModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific vault based on its id
func (m *VaultModel) Get(id int) (*Vault, error) {
	return nil, nil
}

// This will return the 10 most recent created vaults
func (m *VaultModel) Latest() ([]*Vault, error) {
	return nil, nil
}
