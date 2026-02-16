package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// =====================
// Spotify API Response Types
// =====================

type spotifyUserTopArtistsResponse struct {
	Items []struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		Genres     []string `json:"genres"`
		Popularity int      `json:"popularity"`
		Followers  struct {
			Total int `json:"total"`
		} `json:"followers"`
		Images []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
}

type spotifyTopTracksResponse struct {
	Tracks []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		DurationMs int    `json:"duration_ms"`
		PreviewURL string `json:"preview_url"`
	} `json:"tracks"`
}

// =====================
// HTTP Helper
// =====================

func spotifyGet(token, endpoint string, target interface{}) error {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/"+endpoint, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("spotify API error (%d): %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// =====================
// Main Function: GetTopArtists
// Uses GET /me/top/artists - returns the logged-in user's top artists
// based on their actual listening history. Requires user-top-read scope.
// =====================

func GetTopArtists(token string) ([]models.Artist, error) {
	// Fetch user's top 50 artists (medium_term = last ~6 months)
	var result spotifyUserTopArtistsResponse
	if err := spotifyGet(token, "me/top/artists?limit=50&time_range=medium_term", &result); err != nil {
		return nil, fmt.Errorf("fetching user top artists: %w", err)
	}

	var artists []models.Artist
	for _, a := range result.Items {
		image := ""
		if len(a.Images) > 0 {
			image = a.Images[0].URL
		}

		genres := make([]string, len(a.Genres))
		for i, g := range a.Genres {
			genres[i] = capitalizeGenre(g)
		}

		artist := models.Artist{
			ID:         a.ID,
			Name:       a.Name,
			Image:      image,
			Genres:     genres,
			Popularity: a.Popularity,
			Followers:  a.Followers.Total,
		}

		// Fetch top 3 tracks - non-fatal if this fails
		tracks, err := getArtistTopTracks(token, a.ID)
		if err == nil {
			artist.TopTracks = tracks
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

// getArtistTopTracks fetches top 3 tracks for an artist
func getArtistTopTracks(token, artistID string) ([]models.Track, error) {
	var result spotifyTopTracksResponse
	if err := spotifyGet(token, fmt.Sprintf("artists/%s/top-tracks?market=GB", artistID), &result); err != nil {
		return nil, err
	}

	limit := 3
	if len(result.Tracks) < limit {
		limit = len(result.Tracks)
	}

	var tracks []models.Track
	for _, t := range result.Tracks[:limit] {
		tracks = append(tracks, models.Track{
			ID:         t.ID,
			Name:       t.Name,
			Duration:   models.FormatDuration(t.DurationMs),
			PreviewURL: t.PreviewURL,
		})
	}

	return tracks, nil
}

// capitalizeGenre formats Spotify's lowercase genre strings
// e.g. "hip hop" → "Hip Hop", "pop" → "Pop"
func capitalizeGenre(g string) string {
	words := strings.Fields(g)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
