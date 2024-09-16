package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// The routes() method returns a new servemux containing application routes
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Handler function which wraps notFound() helper, and then assign it as the custom handler for
	// 404 Not Found response.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/clip/view/:id", dynamic.ThenFunc(app.clipView))
	router.Handler(http.MethodGet, "/clip/create", dynamic.ThenFunc(app.clipCreate))
	router.Handler(http.MethodPost, "/clip/create", dynamic.ThenFunc(app.clipCreatePost))

	standard := alice.New(app.revocerPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
