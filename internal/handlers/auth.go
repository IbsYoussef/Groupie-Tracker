package handlers

import "net/http"

// LoginHandler serves the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}

// RegisterHandler serves the register page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", nil)
}
