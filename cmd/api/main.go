package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// === DATABASE INITIALIZATION ===
	if err := database.Initialize(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer database.Close()

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

	// Create new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	// Serve CSS, JS, and Images
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// ===== PUBLIC ROUTES =====
	// Landing page
	mux.HandleFunc("GET /", handlers.LandingHandler)

	// Auth pages
	mux.HandleFunc("GET /login", handlers.LoginHandler)
	mux.HandleFunc("GET /register", handlers.RegisterHandler)
	mux.HandleFunc("POST /register", handlers.RegisterUserHandler)

	// === PROTECTED ROUTES ===
	mux.HandleFunc("GET /discover", handlers.DiscoverHandler)
	mux.HandleFunc("GET /logout", handlers.LogoutHandler)

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	// ===== SERVER STARTUP =====
	log.Printf("🚀 Server starting on port %s in %s mode", port, env)
	log.Printf("📍 Visit: http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
