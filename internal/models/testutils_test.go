package models

import (
	"database/sql"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	// Establish a sql.DB connection pool for test database.
	db, err := sql.Open(
		"postgres",
		"host=localhot port=5433 user=testing_web password=testing_web dbname=test_clipvault sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}

	// Read the setup SQL script from file and execute the statements
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Use the t.Cleanup() to register a function *which will automatically be called by Go
	// when the current test (or sub-test) which call newTestDB() has finished*. In this function
	// we read and execute the teardown script, and close the database connection pool
	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	// Return the database connection pool
	return db
}
