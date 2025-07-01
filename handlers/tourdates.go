package handlers

import (
	"groupie-tracker/packages/models"
	"net/http"
)

func TourDatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tour-dates" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	Tpl.ExecuteTemplate(w, "tour-dates.html", models.API)
}
