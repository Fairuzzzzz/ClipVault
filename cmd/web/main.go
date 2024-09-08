package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// Create a logger for writing information message
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize instance of application struct, containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/clip/view", app.clipView)
	mux.HandleFunc("/clip/create", app.clipCreate)

	// Initialize a http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// The value returned form flag.String() function is a pointer to the flag value
	// not the value itself
	infoLog.Printf("Starting server on %s", *addr) // Information message
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // Error message
}
