package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection pool
var DB *sql.DB

// Connect establishes a connection to PostgreSQL using provided config
func Connect(host, port, user, password, dbName, sslMode string) (*sql.DB, error) {
	// Build connection string (DSN - Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL")
	return db, nil
}

// Initialize connects to the database and sets the global DB variable
func Initialize(host, port, user, password, dbName, sslMode string) error {
	log.Printf("🔌 Connecting to PostgreSQL at %s:%s...", host, port)

	db, err := Connect(host, port, user, password, dbName, sslMode)
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
