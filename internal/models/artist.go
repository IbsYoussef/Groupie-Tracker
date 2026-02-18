package models

import (
	"database/sql"
	"fmt"
)

// Artist represents a Spotify artist with enriched data for display
type Artist struct {
	ID            string
	Name          string
	Image         string
	Genres        []string
	Popularity    int
	Followers     int
	ChartRank     int     // Last.fm chart position
	Playcount     int     // Last.fm playcount
	PopularAlbums []Album // Top albums from Spotify
	TopTracks     []Track
}

// Album represents a Spotify album
type Album struct {
	ID          string
	Name        string
	Image       string
	ReleaseDate string
}

// Track represents a Spotify track with display-ready duration
type Track struct {
	ID         string
	Name       string
	Duration   string
	PreviewURL string
}

// FormatDuration converts Spotify's durations_ms to "m:ss" format
func FormatDuration(ms int) string {
	totalSeconds := ms / 1000
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%d:02d", minutes, seconds)
}

// SaveSpotifyToken stores the Spotify access token against a user.
// Called on every login so the token stays fresh.
func SaveSpotifyToken(db *sql.DB, userID, token string) error {
	query := `UPDATE users SET spotify_access_token = $1 WHERE id = $2`
	_, err := db.Exec(query, token, userID)
	return err
}

// GetSpotifyTokenByUserID retrieves the stored Spotify access token for a user.
// Returns an empty string if no token is stored.
func GetSpotifyTokenByUserID(db *sql.DB, userID string) (string, error) {
	var token string
	query := `SELECT COALESCE(spotify_access_token, '') FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&token)
	if err != nil {
		return "", fmt.Errorf("fetching spotify token: %w", err)
	}
	return token, nil
}
