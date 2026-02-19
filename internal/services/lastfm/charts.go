package lastfm

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// =====================
// Last.fm API Response Types
// =====================

type lastfmTopArtistsResponse struct {
	Artists struct {
		Artist []struct {
			Name      string `json:"name"`
			Playcount string `json:"playcount"`
			Listeners string `json:"listeners"`
			MBID      string `json:"mbid"`
			URL       string `json:"url"`
			Attr      struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
		} `json:"artist"`
	} `json:"artists"`
}

// ChartArtist represents an artist from Last.fm charts
type ChartArtist struct {
	Name      string
	Rank      int
	Playcount int
	Listeners int
	ImageURL  string
}

// GetTopArtists fetches the top 50 artists from Last.fm global charts
func GetTopArtists(apiKey string) ([]ChartArtist, error) {
	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&limit=50&api_key=%s&format=json",
		apiKey,
	)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching Last.fm charts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Last.fm API error (%d): %s", resp.StatusCode, string(body))
	}

	var result lastfmTopArtistsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("parsing Last.fm response: %w", err)
	}

	var artists []ChartArtist
	for i, a := range result.Artists.Artist {
		// Parse rank (already a string number)
		rank := i + 1 // Use index + 1 as fallback
		fmt.Sscanf(a.Attr.Rank, "%d", &rank)

		// Parse playcount and listeners
		playcount := 0
		listeners := 0
		fmt.Sscanf(a.Playcount, "%d", &playcount)
		fmt.Sscanf(a.Listeners, "%d", &listeners)

		// Get largest image
		imageURL := ""
		for _, img := range a.Image {
			if img.Size == "extralarge" || img.Size == "mega" {
				imageURL = img.Text
				break
			}
		}

		// Fallback to any available image if extralarge/mega not found
		if imageURL == "" && len(a.Image) > 0 {
			// Try large, medium, small in order
			for _, img := range a.Image {
				if img.Size == "large" || img.Size == "medium" || img.Size == "small" {
					imageURL = img.Text
					break
				}
			}
		}

		artist := ChartArtist{
			Name:      a.Name,
			Rank:      rank,
			Playcount: playcount,
			Listeners: listeners,
			ImageURL:  imageURL,
		}

		artists = append(artists, artist)

		// Debug: Log first 5 artists with their image status
		if len(artists) <= 5 {
			if imageURL != "" {
				log.Printf("🖼️  Rank #%d: %s - Image: %.50s...", artist.Rank, artist.Name, imageURL)
			} else {
				log.Printf("⚠️  Rank #%d: %s - NO IMAGE", artist.Rank, artist.Name)
			}
		}
	}

	return artists, nil
}
