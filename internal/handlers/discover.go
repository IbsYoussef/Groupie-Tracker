package handlers

import (
	"log"
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// DiscoverPageData holds template data for the discover page
type DiscoverPageData struct {
	User *models.User
}

// DiscoverHandler serves the discover page (requires authentication)
func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	cfg := config.Load()

	if !models.VerifySessionToken(cfg.SessionSecret, cookie.Value) {
		log.Printf("⚠️  Invalid session token signature")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := models.GetSessionByToken(database.DB, cookie.Value)
	if err != nil || session == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get user to pass to template
	user, err := models.GetUserByID(database.DB, session.UserID)
	if err != nil {
		log.Printf("❌ Error fetching user: %v", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Printf("✅ Discover page loaded for user: %s", user.Username)

	// Use the existing RenderTemplate helper (components now loaded via init())
	RenderTemplate(w, "discover.html", DiscoverPageData{User: user})
}
