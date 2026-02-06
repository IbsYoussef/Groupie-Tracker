package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// DB is the global connection configuration
var DB *sql.DB

// Config holds database connection configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// LoadConfig loads database configuration from environment variables
// Falls back to defaults if not set
func LoadConfig() Config {
	return Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "groupie_user"),
		Password: getEnv("DB_PASSWORD", "groupie_password"),
		DBName:   getEnv("DB_NAME", "groupie_tracker"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// Connect establishes a connection to PostgreSQL
func Connect(cfg Config) (*sql.DB, error) {
	// Build connection string (DSN - Data Source Name)
	// Format: "postgres://user:password@host:port/dbname?sslmode=disable"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Open connection (doesn't actually connect yet, just prepares)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	// MaxOpenConns: Maximum number of open connections to the database
	db.SetMaxOpenConns(25)

	// MaxIdleConns: Maxiumum number off idle connections in the pool
	db.SetMaxIdleConns(5)

	// ConnMaxLifetime: Maximum time a connection can be reused
	db.SetConnMaxLifetime(5 * time.Minute)

	// Actually test the connection with a ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL")
	return db, nil
}

// Initialize connects to the database and sets the global DB variable
func Initialize() error {
	cfg := LoadConfig()

	log.Printf("🔌 Connecting to PostgreSQL at %s:%s...", cfg.Host, cfg.Port)

	db, err := Connect(cfg)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		log.Println("🔌 Closing database connection...")
		return DB.Close()
	}

	return nil
}
