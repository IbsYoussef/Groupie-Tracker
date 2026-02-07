package handlers

import (
	"log"
	"net/http"

	"github.com/IbsYoussef/Groupie-Tracker/internal/database"
	"github.com/IbsYoussef/Groupie-Tracker/internal/models"
)

// LoginHandler serves the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.html", nil)
}

// RegisterHandler serves the register page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "register.html", nil)
}

// RegisterUserHandler processes the registration form
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// Validation
	if username == "" || email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if password != confirmPassword {
		http.Error(w, "Password do not match", http.StatusBadRequest)
		return
	}

	if len(password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	// Create user in database
	user, err := models.CreateUser(database.DB, username, email, password)
	if err != nil {
		if err == models.ErrEmailExists {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		if err == models.ErrUsernameExists {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ User created: %s (%s)", user.Username, user.Email)

	// Create session
	session, err := models.CreateSession(database.DB, user.ID)
	if err != nil {
		log.Printf("Error creating session: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		Expires:  session.ExpiresAt,
		HttpOnly: true,  // Prevent JavaScript access (XSS protection)
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	log.Printf("✅ Session created for user: %s", user.Username)

	// Redirect to discover page
	http.Redirect(w, r, "/discover", http.StatusSeeOther)
}

// LoginUserHandler processes the login form
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validation
	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Authenticate user
	user, err := models.Authenticate(database.DB, email, password)
	if err != nil {
		if err == models.ErrInvalidCredentials {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		log.Printf("Error authenticating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ User authenticated: %s", user.Username)

	// Create session
	session, err := models.CreateSession(database.DB, user.ID)
	if err != nil {
		log.Printf("Error creating session: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Path:     "/",
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	log.Printf("✅ Session created for user: %s", user.Username)

	// Redirect to discover page
	http.Redirect(w, r, "/discover", http.StatusSeeOther)
}

// LogoutHandler logs out the user
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get session token from cookie
	cookie, err := r.Cookie("session_token")
	if err == nil {
		// Delete session from database
		models.DeleteSession(database.DB, cookie.Value)
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
		HttpOnly: true,
	})

	// Redirect to landing page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
