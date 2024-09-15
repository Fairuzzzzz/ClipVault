package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fairuzzzzz/clipvault/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	clips, err := app.clips.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Clips = clips

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) clipView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	clip, err := app.clips.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Clip = clip

	app.render(w, http.StatusOK, "view.html", data)
}

func (app *application) clipCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet..."))
}

func (app *application) clipCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Test"
	content := "Test"
	expires := 7

	id, err := app.clips.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/clip/view/%d", id), http.StatusSeeOther)
}
