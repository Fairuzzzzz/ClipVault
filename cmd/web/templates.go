package main

import "github.com/Fairuzzzzz/clipvault/internal/models"

// Define templateData type to act as the holding structure for any dynamic data that want to pass to HTML templates.
type templateData struct {
	Clip *models.Clip
}
