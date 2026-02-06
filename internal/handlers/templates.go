package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Global template variable - loaded once at startup
var tpl *template.Template

// init runs automatically when the package is imported
func init() {
	var err error

	// Parse all templates from the templates directory
	templatePath := filepath.Join("web", "templates", "*.html")
	tpl, err = template.ParseGlob(templatePath)
	if err != nil {
		log.Fatalf("❌ Error parsing templates: %v", err)
	}

	log.Println("✅ Templates loaded successfully")
}

// RenderTemplate is a helper function to render templates with error handling
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute template
	if err := tpl.ExecuteTemplate(w, tmpl, data); err != nil {
		log.Printf("❌ Error executing template %s: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
