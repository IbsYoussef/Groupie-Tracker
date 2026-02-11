package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IbsYoussef/Groupie-Tracker/internal/config"
	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/handlers"
	"github.com/IbsYoussef/Groupie-Tracker/internal/services/spotify"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables in development
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		}
	}

	// Load configuration
	cfg := config.Load()

	// === DATABASE INITIALIZATION ===
	if err := database.Initialize(
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBUSER,
		cfg.DBPASSWORD,
		cfg.DBNAME,
		cfg.DBSSLMODE,
	); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer func() {
		_ = database.Close()
	}()

	// Create new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// ===== PUBLIC ROUTES =====
	mux.HandleFunc("GET /", handlers.LandingHandler)
	mux.HandleFunc("GET /login", handlers.LoginHandler)
	mux.HandleFunc("GET /register", handlers.RegisterHandler)
	mux.HandleFunc("POST /register", handlers.RegisterUserHandler)
	mux.HandleFunc("POST /login", handlers.LoginUserHandler)

	// OAuth route
	mux.HandleFunc("GET /auth/spotify", spotify.SpotifyLoginHandler)
	mux.HandleFunc("GET /auth/spotify/callback", spotify.SpotifyCallbackHandler)

	// === PROTECTED ROUTES ===
	mux.HandleFunc("GET /discover", handlers.DiscoverHandler)
	mux.HandleFunc("GET /logout", handlers.LogoutHandler)

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	// ===== SERVER STARTUP =====
	log.Printf("🚀 Server starting on port %s in %s mode", cfg.Port, cfg.Env)
	log.Printf("📍 Visit: http://127.0.0.1:%s", cfg.Port)

	if err := http.ListenAndServe("127.0.0.1:"+cfg.Port, mux); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
