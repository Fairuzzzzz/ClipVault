package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Fairuzzzzz/clipvault/internal/models"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	clips    *models.ClipModel
}

func main() {
	// Command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "host=localhost user=web2 password=testweb dbname=clipvault sslmode=disable", "PostgreSQL data source name")

	flag.Parse()

	// Create a logger for writing information message
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initialize instance of application struct, containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		clips:    &models.ClipModel{DB: db},
	}

	// Initialize a http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// The value returned form flag.String() function is a pointer to the flag value
	// not the value itself
	infoLog.Printf("Starting server on %s", *addr) // Information message
	err = srv.ListenAndServe()
	errorLog.Fatal(err) // Error message
}

// The openDB() function wraps sqlx.Connect() and returns a sqlx.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
