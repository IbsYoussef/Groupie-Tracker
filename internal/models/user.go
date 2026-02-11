package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the database
type User struct {
	ID            string
	Username      string
	Email         string
	PasswordHash  sql.NullString // Null for OAuth user
	OAuthProvider sql.NullString
	OAuthID       sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Common errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailExists        = errors.New("email already exists")
	ErrUsernameExists     = errors.New("username already exists")
)

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	// Cost 12 = good balance of security vs speed
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPassword compares a plain-text password with a hash
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, username, email, password string) (*User, error) {
	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// SQL query with RETURNING to get the created user
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, username, email, created_at, updated_at
	`

	user := &User{}
	err = db.QueryRow(query, username, email, hashedPassword).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		// Check for unique constraint violations
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			return nil, ErrEmailExists
		}
		if err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"` {
			return nil, ErrUsernameExists
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &User{}
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(db *sql.DB, id string) (*User, error) {
	query := `
		SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &User{}
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Authenticate checks email and password, returns user if valid
func Authenticate(db *sql.DB, email, password string) (*User, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Check if password_hash is NULL (OAuth user)
	if !user.PasswordHash.Valid {
		return nil, ErrInvalidCredentials
	}

	// Compare password with hash
	if err := CheckPassword(password, user.PasswordHash.String); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// GetUserByOAuth retrieves a user by OAuth provider and ID
func GetUserByOAuth(db *sql.DB, provider, oauthID string) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, oauth_provider, oauth_id, created_at, updated_at
		FROM users
		WHERE oauth_provider = $1 AND oauth_id = $2
	`

	user := &User{}
	err := db.QueryRow(query, provider, oauthID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.OAuthProvider,
		&user.OAuthID,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateOAuthUser creates a new user from OAuth provider
func CreateOAuthUser(db *sql.DB, username, email, provider, oauthID string) (*User, error) {
	query := `
		INSERT INTO users (username, email, oauth_provider, oauth_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, password_hash, oauth_provider, oauth_id, created_at, updated_at
	`

	user := &User{}
	err := db.QueryRow(query, username, email, provider, oauthID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.OAuthProvider,
		&user.OAuthID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		// Check for duplicate email
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			return nil, ErrEmailExists
		}
		// Check for duplicate username
		if err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"` {
			return nil, ErrUsernameExists
		}
		return nil, err
	}
	return user, nil
}
