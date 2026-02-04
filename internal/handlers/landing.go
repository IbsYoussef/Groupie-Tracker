package handlers

import (
	"net/http"
)

// LandingHandler serves the landing page
func LandingHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}
