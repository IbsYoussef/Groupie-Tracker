package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
// SIMPLE VERSION - Last.fm Only
// =====================

// GetTopArtistsSimple fetches top 50 artists from Last.fm
// Uses Spotify for images only (if token provided)
func GetTopArtistsSimple(lastfmAPIKey string) ([]models.Artist, error) {
	// Get Last.fm chart
	chartArtists, err := lastfm.GetTopArtists(lastfmAPIKey)
	if err != nil {
		return nil, fmt.Errorf("fetching Last.fm charts: %w", err)
	}

	log.Printf("📊 Last.fm returned %d artists", len(chartArtists))

	// Convert Last.fm artists to our model
	// Note: Using Last.fm images which may be blocked by some networks
	// For production, consider implementing Spotify image enrichment
	artists := make([]models.Artist, 0, len(chartArtists))
	for _, ca := range chartArtists {
		artists = append(artists, models.Artist{
			Name:          ca.Name,
			Image:         ca.ImageURL, // Last.fm CDN may be blocked
			ChartRank:     ca.Rank,
			Playcount:     ca.Playcount,
			Followers:     ca.Listeners,
			Genres:        []string{},
			PopularAlbums: []models.Album{},
		})
	}

	log.Printf("✅ Loaded %d artists from Last.fm", len(artists))
	log.Printf("⚠️  Using Last.fm images - if images don't display, the CDN may be blocked")

	return artists, nil
}

// =====================
// ENRICHED VERSION - Last.fm + Spotify (TODO)
// =====================

// GetTopArtistsWithSpotify fetches from Last.fm and enriches with Spotify images
// Falls back to Last.fm data if Spotify enrichment fails
func GetTopArtistsWithSpotify(spotifyToken, lastfmAPIKey string) ([]models.Artist, error) {
	// Get Last.fm chart
	chartArtists, err := lastfm.GetTopArtists(lastfmAPIKey)
	if err != nil {
		return nil, fmt.Errorf("fetching Last.fm charts: %w", err)
	}

	log.Printf("📊 Last.fm returned %d artists", len(chartArtists))

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		artists []models.Artist
	)

	// Use buffered channel to limit concurrent requests
	semaphore := make(chan struct{}, 10)

	for _, chartArtist := range chartArtists {
		wg.Add(1)

		go func(ca lastfm.ChartArtist) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Start with Last.fm data (guaranteed to exist)
			artist := models.Artist{
				Name:          ca.Name,
				Image:         ca.ImageURL, // Fallback to Last.fm image
				ChartRank:     ca.Rank,
				Playcount:     ca.Playcount,
				Followers:     ca.Listeners,
				Genres:        []string{},
				PopularAlbums: []models.Album{},
			}

			// Try to get Spotify image (simple search, no complex validation)
			if spotifyToken != "" {
				if spotifyArtist, err := searchArtistForImage(spotifyToken, ca.Name); err == nil && spotifyArtist != nil {
					// Use Spotify image (higher quality and not blocked)
					if spotifyArtist.Image != "" {
						artist.Image = spotifyArtist.Image
					}
					// Bonus: get genres and ID if available
					artist.ID = spotifyArtist.ID
					artist.Genres = spotifyArtist.Genres
				}
			}

			mu.Lock()
			artists = append(artists, artist)
			mu.Unlock()
		}(chartArtist)
	}

	wg.Wait()

	// Sort by chart rank to maintain order
	sortArtistsByRank(artists)

	log.Printf("✅ Loaded %d artists (with Spotify images)", len(artists))

	return artists, nil
}

// searchArtistForImage does a simple Spotify search just to get the image
// Uses basic name validation to avoid completely wrong matches
func searchArtistForImage(token, name string) (*models.Artist, error) {
	endpoint := fmt.Sprintf(
		"search?q=%s&type=artist&limit=3",
		url.QueryEscape(name),
	)

	var result spotifySearchResponse
	if err := spotifyGet(token, endpoint, &result); err != nil {
		return nil, err
	}

	if len(result.Artists.Items) == 0 {
		return nil, nil
	}

	// Simple validation: check if artist name is somewhat similar
	searchLower := strings.ToLower(strings.TrimSpace(name))

	// Try to find a match in top 3 results
	for _, candidate := range result.Artists.Items {
		candidateLower := strings.ToLower(strings.TrimSpace(candidate.Name))

		// Check if names match (case-insensitive, ignoring "The" prefix)
		searchClean := strings.TrimPrefix(searchLower, "the ")
		candidateClean := strings.TrimPrefix(candidateLower, "the ")

		// Match if: exact match, one contains the other, or very similar
		if searchClean == candidateClean ||
			strings.Contains(candidateClean, searchClean) ||
			strings.Contains(searchClean, candidateClean) {

			// Get highest quality image
			image := ""
			if len(candidate.Images) > 0 {
				image = candidate.Images[0].URL
			}

			// Capitalize genres
			genres := make([]string, len(candidate.Genres))
			for i, g := range candidate.Genres {
				genres[i] = capitalizeGenre(g)
			}

			return &models.Artist{
				ID:     candidate.ID,
				Image:  image,
				Genres: genres,
			}, nil
		}
	}

	// No good match found - return nil (will use Last.fm fallback)
	return nil, nil
}

// sortArtistsByRank sorts artists by their Last.fm chart position
func sortArtistsByRank(artists []models.Artist) {
	// Simple bubble sort since we have <50 artists
	for i := 0; i < len(artists); i++ {
		for j := i + 1; j < len(artists); j++ {
			if artists[i].ChartRank > artists[j].ChartRank {
				artists[i], artists[j] = artists[j], artists[i]
			}
		}
	}
}

// =====================
// Helper Functions (for future Spotify enrichment)
// =====================

// getArtistAlbums fetches top 3 albums for an artist
func getArtistAlbums(token, artistID string) ([]models.Album, error) {
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
func capitalizeGenre(g string) string {
	words := strings.Fields(g)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
