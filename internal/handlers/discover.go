package handlers

import (
	"log"
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// DiscoverHandler serves the discover page (requires authentication)
func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	// Get session token from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Load config for session secret
	cfg := config.Load()

	// Verify token signature
	if !models.VerifySessionToken(cfg.SessionSecret, cookie.Value) {
		log.Printf("⚠️  Invalid session token signature")
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
