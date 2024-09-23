package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModelInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (int, error)
	Exists(id int) (bool, error)
	Get(id int) (*User, error)
}

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
	var id int
	var hashedPassowrd []byte

	stmt := `SELECT id, hashed_password FROM users WHERE email = $1`

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassowrd)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// Check whether the hashed password and plain-text password provided match
	// If the don't, return the ErrInvalidCredentials error
	err = bcrypt.CompareHashAndPassword(hashedPassowrd, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}

// Exists method to check if a user exists with the specific ID
func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool

	stmt := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`

	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

func (m *UserModel) Get(id int) (*User, error) {
	var user User

	stmt := `SELECT id, name, email, created FROM users WHERE id = $1`

	err := m.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Name, &user.Email, &user.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &user, nil
}
