package handlers

import (
	"net/http"
)

// LandingHandler serves the landing page
func LandingHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}
