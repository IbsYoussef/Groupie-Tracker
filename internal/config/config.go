package config

import "os"

// Config holds all application configuration
type Config struct {
	// Database
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
	DBSSLMODE  string

	// Server
	Port string
	Env  string

	// Session
	SessionSecret string

	// APIs (future implementations)
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURI  string
	OpenAIAPIKey        string
}

// Load returns a new Config with values from environment variables
func Load() *Config {
	return &Config{
		// Database
		DBHOST:     getEnv("DB_HOST", "localhost"),
		DBPORT:     getEnv("DB_PORT", "5432"),
		DBUSER:     getEnv("DB_USER", "groupie_user"),
		DBPASSWORD: getEnv("DB_PASSWORD", "groupie_password"),
		DBNAME:     getEnv("DB_NAME", "groupie_tracker"),
		DBSSLMODE:  getEnv("DB_SSLMODE", "disable"),

		// Server
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "development"),

		// Session
		SessionSecret: getEnv("SESSION_SECRET", ""),

		// APIs
		SpotifyClientID:     getEnv("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyRedirectURI:  getEnv("SPOTIFY_REDIRECT_URI", ""),
		OpenAIAPIKey:        getEnv("OPENAI_API_KEY", ""),
	}
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
