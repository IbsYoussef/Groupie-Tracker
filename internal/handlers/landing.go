package handlers

import (
	"net/http"
)

// LandingHandler serves the landing page
func LandingHandler(w http.ResponseWriter, r *http.Request) {
	// Data to pass to template
	data := struct {
		Title       string
		Description string
		Year        int
	}{
		Title:       "Groupie Tracker - Discover Your Music Universe",
		Description: "Find artists, explore concerts, powered by Spotify",
		Year:        2025,
	}

	// Use the centralized renderTemplate helper
	renderTemplate(w, "landing.html", data)
}
