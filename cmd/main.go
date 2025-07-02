package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Local fallback port
	}

	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	fmt.Println("👉 Use Ctrl+C to stop the server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
