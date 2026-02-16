package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// SpotifyLoginHandler redirects user to Spotify authorization page
func SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()

	// Build Spotify Authorization URL
	authURL := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		cfg.SpotifyClientID,
		url.QueryEscape(cfg.SpotifyRedirectURI),
		url.QueryEscape("user-read-email user-read-private user-top-read"),
	)

	log.Println("🎵 Redirecting to Spotify login...")
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// SpotifyCallbackHandler handles the callback from Spotify
func SpotifyCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("🔔 Callback hit! Full URL: %s", r.URL.String())
	log.Printf("🔔 Code: %s", r.URL.Query().Get("code"))
	if errParam := r.URL.Query().Get("error"); errParam != "" {
		log.Printf("❌ Spotify auth error: %s", errParam)
	}

	cfg := config.Load()

	// Get authorization code from query params
	code := r.URL.Query().Get("code")
	if code == "" {
		log.Println("❌ No authorization code from Spotify")
		http.Error(w, "Authorization failed", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Received authorization code from Spotify")

	// Exchange code for access token
	tokenURL := "https://accounts.spotify.com/api/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", cfg.SpotifyRedirectURI)
	data.Set("client_id", cfg.SpotifyClientID)
	data.Set("client_secret", cfg.SpotifyClientSecret)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		log.Printf("❌ Error changing code for token: %v", err)
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("❌ Error reading response body: %v", err)
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.Printf("❌ Error parsing token response: %v", err)
		http.Error(w, "Failed to parse token", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Got access token from Spotify")

	// Get user profile from Spotify
	profileReq, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	profileReq.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)

	client := &http.Client{}
	profileResp, err := client.Do(profileReq)
	if err != nil {
		log.Printf("❌ Error fetching Spotify profile: %v", err)
		http.Error(w, "Failed to get profile", http.StatusInternalServerError)
		return
	}
	defer profileResp.Body.Close()

	profileBody, _ := io.ReadAll(profileResp.Body)

	var profile struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	}

	if err := json.Unmarshal(profileBody, &profile); err != nil {
		log.Printf("❌ Error parsing Spotify profile: %v", err)
		http.Error(w, "Failed to parse profile", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Got Spotify profile: %s (%s)", profile.DisplayName, profile.Email)

	// Check if user exists
	user, err := models.GetUserByOAuth(database.DB, "spotify", profile.ID)
	if err != nil && err != models.ErrUserNotFound {
		log.Printf("❌ Database error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create a new user if they don't exist
	if user == nil {
		log.Printf("🆕 Creating new Spotify user: %s", profile.Email)
		user, err = models.CreateOAuthUser(database.DB, profile.DisplayName, profile.Email, "spotify", profile.ID)
		if err != nil {
			log.Printf("❌ Error creating OAuth user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
		log.Printf("✅ Spotify user created: %s", user.Username)
	} else {
		log.Printf("✅ Existing Spotify user: %s", user.Username)
	}

	// ── Store access token against user for later API calls ──
	if err := models.SaveSpotifyToken(database.DB, user.ID, tokenResponse.AccessToken); err != nil {
		log.Printf("⚠️  Could not save Spotify token: %v", err)
		// Non-fatal - continue to login
	}

	// Create session
	session, err := models.CreateSession(database.DB, user.ID, false, cfg.SessionSecret)
	if err != nil {
		log.Printf("❌ Error creating session: %v", err)
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	log.Printf("✅ Session created for Spotify user: %s (expires: %s)", user.Username, session.ExpiresAt.Format("2006-01-02 15:04:05"))

	// Redirect to discover page
	http.Redirect(w, r, "/discover", http.StatusSeeOther)
}
