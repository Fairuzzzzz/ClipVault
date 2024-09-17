package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Define a UserModel type which wraps a database connection pool
type UserModel struct {
	DB *sql.DB
}

// Inser method to add a new record to the "users" table
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES ($1, $2, $3, CURRENT_TIMESTAMP AT TIME ZONE 'UTC')`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassowrd))
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" && pqErr.Constraint == "users_email_key" {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

// Authenticate method to verify whether a user exists with the provided email
// address and passowrd.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists method to check if a user exists with the specific ID
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
