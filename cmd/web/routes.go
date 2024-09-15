package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a new servemux containing application routes
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/clip/view", app.clipView)
	mux.HandleFunc("/clip/create", app.clipCreate)

	standard := alice.New(app.revocerPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
