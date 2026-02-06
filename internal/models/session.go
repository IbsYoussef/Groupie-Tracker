package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"
)

// Session represents a user session in the database
type Session struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

// Session duration
const (
	SessionDuration = 7 * 24 * time.Hour // 7 days
)

// GenerateSessionToken creates a secure random token
func GenerateSessionToken() (string, error) {
	// Generate 32 random bytes
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode to base64 (URL-safe)
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// CreateSession creates a new session for a user
func CreateSession(db *sql.DB, userID string) (*Session, error) {
	// Generate secure token
	token, err := GenerateSessionToken()
	if err != nil {
		return nil, err
	}

	// Set expiration
	expiresAt := time.Now().Add(SessionDuration)

	// Insert into database
	query := `
		INSERT INTO sessions (user_id, token, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, token, expires_at, created_at
	`

	session := &Session{}
	err = db.QueryRow(query, userID, token, expiresAt).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.ExpiresAt,
		&session.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// GetSessionsByToken retrieves a session by token
func GetSessionByToken(db *sql.DB, token string) (*Session, error) {
	query := `
		SELECT id, user_id, token, expires_at, created_at
		FROM sessions
		WHERE token = $1 AND expires_at > NOW()
	`

	session := &Session{}
	err := db.QueryRow(query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.ExpiresAt,
		&session.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Session not found or expired
	}

	if err != nil {
		return nil, err
	}

	return session, nil
}

// Delete session removes a session (logout)
func DeleteSession(db *sql.DB, token string) error {
	query := `DELETE FROM sessions WHERE token = $1`
	_, err := db.Exec(query, token)
	return err
}

// CleanupExpiredSessions removes all expired sessions (run periodically)
func CleanupExpiredSessions(db *sql.DB) error {
	query := `DELETE FROM sessions WHERE expires_at < NOW()`
	_, err := db.Exec(query)
	return err
}
