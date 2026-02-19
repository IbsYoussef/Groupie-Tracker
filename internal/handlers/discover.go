package handlers

import (
	"log"
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
	"github.com/IbsYoussef/Groupie-Tracker/internal/services/spotify"
)

// DiscoverPageData holds all data passed to the discover template
type DiscoverPageData struct {
	User    *models.User
	Artists []models.Artist
	Error   string
}

// DiscoverHandler serves the discover page (requires authentication)
func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	// ── Auth check ──────────────────────────────────────────
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

	user, err := models.GetUserByID(database.DB, session.UserID)
	if err != nil {
		log.Printf("❌ Error fetching user: %v", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// ── Fetch top artists from Last.fm + Spotify images ──────────────────────
	// Get Spotify token for image enrichment (optional - works without it too)
	spotifyToken, _ := models.GetSpotifyTokenByUserID(database.DB, session.UserID)

	artists, err := spotify.GetTopArtistsWithSpotify(spotifyToken, cfg.LastFMAPIKey)
	if err != nil {
		log.Printf("❌ Error fetching top artists: %v", err)
		RenderTemplate(w, "discover.html", DiscoverPageData{
			User:  user,
			Error: "Unable to load artists right now. Please try again later.",
		})
		return
	}

	log.Printf("✅ Discover page loaded for user: %s (%d artists)", user.Username, len(artists))

	// Debug: Check first 3 artists
	for i := 0; i < min(3, len(artists)); i++ {
		log.Printf("🔍 Artist #%d: Name=%s, Image=%s, Rank=%d",
			i+1, artists[i].Name, artists[i].Image, artists[i].ChartRank)
	}

	RenderTemplate(w, "discover.html", DiscoverPageData{
		User:    user,
		Artists: artists,
	})
}
