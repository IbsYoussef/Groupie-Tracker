package handlers

import (
	"log"
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
	"github.com/IbsYoussef/Groupie-Tracker/internal/services/spotify"
)

// DiscoverPageData holds template data for the discover page
type DiscoverPageData struct {
	User    *models.User
	Artists []models.Artist
	Error   string
}

// DiscoverHandler serves the discover page (requires authentication)
func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	// ── Auth check
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

	// ── Fetch top artists from Spotify ──────────────────────
	// Use the user's OAuth token (already stored in DB from login)
	spotifyToken, err := models.GetSpotifyTokenByUserID(database.DB, session.UserID)
	if err != nil || spotifyToken == "" {
		log.Printf("❌ No Spotify token for user: %v", err)
		RenderTemplate(w, "discover.html", DiscoverPageData{
			User:  user,
			Error: "Unable to load artists. Please log in with Spotify again.",
		})
		return
	}

	artists, err := spotify.GetTopArtists(spotifyToken)
	if err != nil {
		log.Printf("❌ Error fetching top artists: %v", err)
		// Render page with error state rather than crashing
		RenderTemplate(w, "discover.html", DiscoverPageData{
			User:  user,
			Error: "Unable to load artists right now. Please try again later.",
		})
		return
	}

	log.Printf("✅ Discover page loaded for user: %s", user.Username)

	// Use the existing RenderTemplate helper (components now loaded via init())
	RenderTemplate(w, "discover.html", DiscoverPageData{
		User:    user,
		Artists: artists,
	})
}
