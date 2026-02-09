package models

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
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

// SessionDuration Session duration
const (
	SessionDuration           = 7 * 24 * time.Hour  // 7 days (default)
	RememberMeSessionDuration = 30 * 24 * time.Hour // 30 days (remember me)
)

// GenerateSessionToken creates a secure random token with HMAC signature
func GenerateSessionToken(secret string) (string, error) {
	// Generate 32 random bytes
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode to base64 (URL-safe)
	tokenBase := base64.URLEncoding.EncodeToString(bytes)

	// If secret is provided, add HMAC signature
	if secret != "" {
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(tokenBase))
		signature := hex.EncodeToString(h.Sum(nil))

		// Token format: base.signature
		return fmt.Sprintf("%s.%s", tokenBase, signature), nil
	}

	// Fallback: just return random token (for backward compatability)
	return tokenBase, nil
}

// VerifySessionToken checks if token signature is valid
func VerifySessionToken(secret, token string) bool {
	// If no secret, skip verification (backward compatible)
	if secret == "" {
		return true
	}

	// Split token into base and signature
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false
	}

	tokenBase := parts[0]
	providedSignature := parts[1]

	// Compute expected signature
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(tokenBase))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures (constant-time to prevent timing attacks)
	return hmac.Equal([]byte(expectedSignature), []byte(providedSignature))
}

// CreateSession creates a new session for a user
func CreateSession(db *sql.DB, userID string, rememberMe bool, secret string) (*Session, error) {
	// Generate secure token with signature
	token, err := GenerateSessionToken(secret)
	if err != nil {
		return nil, err
	}

	// Set expiration based on rememberMe flag
	var expiresAt time.Time
	if rememberMe {
		expiresAt = time.Now().Add(RememberMeSessionDuration)
	} else {
		expiresAt = time.Now().Add(SessionDuration)
	}

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

// GetSessionByToken retrieves a session by token
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

// DeleteSession removes a session (logout)
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
