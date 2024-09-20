package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/Fairuzzzzz/clipvault/internal/models"
	"github.com/Fairuzzzzz/clipvault/ui"
)

// Define templateData type to act as the holding structure for any dynamic data that want to pass to HTML templates.
type templateData struct {
	CurrentYear     int
	Clip            *models.Clip
	Clips           []*models.Clip
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a map to act as the cache.
	cache := map[string]*template.Template{}

	// Use fs.Glob() to get a slice of all filepaths in the ui.Files embedded
	// filesystem which match the pattern 'html/pages/*.html'. This essentially
	// gies a slice of all the 'page' templates for the application
	pages, err := fs.Glob(ui.Files, "./ui/html/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one
	for _, page := range pages {

		// Extract the file name (like 'home.html') from the full filepath
		// and assign it to the name variable
		name := filepath.Base(page)

		// A slice containing the filepath patterns for the templates want to parse
		patters := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patters...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
