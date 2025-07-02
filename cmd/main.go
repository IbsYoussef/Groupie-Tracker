package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
	"groupie-tracker/internal"
)

func main() {
	err := internal.LoadData()
	if err != nil {
		log.Fatalf("Failed to load API data: %v", err)
	}

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/tour-dates", handlers.TourDatesHandler)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("🚀 Server running at http://localhost:8080")
	fmt.Println("👉 Use Ctrl+C to stop the server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
