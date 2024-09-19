package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Fairuzzzzz/clipvault/internal/models"
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

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one
	for _, page := range pages {

		// Extract the file name (like 'home.html') from the full filepath
		// and assign it to the name variable
		name := filepath.Base(page)

		// Parse the base template file into a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name (like 'home.html') of the page as the key
		cache[name] = ts
	}
	return cache, nil
}
