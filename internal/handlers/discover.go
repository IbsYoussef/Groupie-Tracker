package handlers

import (
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// DiscoverHandler serves the discover page (requires authentication)
func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	// Get session token from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// No cookie = not logged in
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Verify session exists and is valid
	session, err := models.GetSessionByToken(database.DB, cookie.Value)
	if err != nil || session == nil {
		// Invalid/expired session
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// TODO: Get user data and pass to template
	// For now, just render the template
	RenderTemplate(w, "discover.html", nil)
}
