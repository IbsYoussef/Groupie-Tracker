package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables in development
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Simple handler for testing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>🎵 Groupie Tracker v2</h1>")
		fmt.Fprintf(w, "<p>Coming Soon! This is the new and improved version.</p>")
		fmt.Fprintf(w, "<p>Environment: %s</p>", env)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Printf("🚀 Server starting on port %s in %s mode", port, env)
	log.Printf("📍 Visit: http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
