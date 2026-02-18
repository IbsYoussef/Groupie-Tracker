package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
	"github.com/IbsYoussef/Groupie-Tracker/internal/services/lastfm"
)

// =====================
// Spotify API Response Types
// =====================

type spotifySearchResponse struct {
	Artists struct {
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
	} `json:"artists"`
}

type spotifyAlbumsResponse struct {
	Items []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`
		Images      []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
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
// 1. Fetches Last.fm top 50 global chart
// 2. Enriches each artist with Spotify metadata (search by name)
// 3. Gets top 3 albums for each artist
// =====================

func GetTopArtists(spotifyToken, lastfmAPIKey string) ([]models.Artist, error) {
	// Step 1: Get Last.fm chart
	chartArtists, err := lastfm.GetTopArtists(lastfmAPIKey)
	if err != nil {
		return nil, fmt.Errorf("fetching Last.fm charts: %w", err)
	}

	var artists []models.Artist
	seenIDs := make(map[string]bool) // Track Spotify IDs

	// Step 2 & 3: Enrich with Spotify data
	for _, chartArtist := range chartArtists {
		// Search Spotify for this artist
		artist, err := searchArtist(spotifyToken, chartArtist.Name)
		if err != nil || artist == nil {
			// Skip if not found on Spotify - non-fatal
			continue
		}

		// Skip duplicates (same Spotify ID)
		if seenIDs[artist.ID] {
			continue
		}
		seenIDs[artist.ID] = true

		// Add Last.fm chart data
		artist.ChartRank = chartArtist.Rank
		artist.Playcount = chartArtist.Playcount

		// Get albums from Spotify
		albums, err := getArtistAlbums(spotifyToken, artist.ID)
		if err == nil {
			artist.PopularAlbums = albums
		}

		artists = append(artists, *artist)
	}

	return artists, nil
}

// searchArtist searches Spotify for an artist by name and returns enriched data
func searchArtist(token, name string) (*models.Artist, error) {
	endpoint := fmt.Sprintf("search?q=%s&type=artist&limit=1",
		url.QueryEscape(name),
	)

	var result spotifySearchResponse
	if err := spotifyGet(token, endpoint, &result); err != nil {
		return nil, err
	}

	if len(result.Artists.Items) == 0 {
		return nil, nil
	}

	a := result.Artists.Items[0]

	// Get the highest quality image
	image := ""
	if len(a.Images) > 0 {
		image = a.Images[0].URL
	}

	// Capitalise genres
	genres := make([]string, len(a.Genres))
	for i, g := range a.Genres {
		genres[i] = capitalizeGenre(g)
	}

	return &models.Artist{
		ID:         a.ID,
		Name:       a.Name,
		Image:      image,
		Genres:     genres,
		Popularity: a.Popularity,
		Followers:  a.Followers.Total,
	}, nil
}

// getArtistAlbums fetches top 3 albums for an artist
func getArtistAlbums(token, artistID string) ([]models.Album, error) {
	// Get albums, sorted by release date (newest first)
	endpoint := fmt.Sprintf(
		"artists/%s/albums?include_groups=album&market=US&limit=3",
		artistID,
	)

	var result spotifyAlbumsResponse
	if err := spotifyGet(token, endpoint, &result); err != nil {
		return nil, err
	}

	var albums []models.Album
	for _, a := range result.Items {
		image := ""
		if len(a.Images) > 0 {
			image = a.Images[0].URL
		}

		// Extract just the year from release date
		releaseYear := a.ReleaseDate
		if len(releaseYear) >= 4 {
			releaseYear = releaseYear[:4]
		}

		albums = append(albums, models.Album{
			ID:          a.ID,
			Name:        a.Name,
			Image:       image,
			ReleaseDate: releaseYear,
		})
	}

	return albums, nil
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
