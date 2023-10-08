package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist/", handlers.ArtistPage)
	http.HandleFunc("/search/", handlers.Search)
	http.HandleFunc("/filters/", handlers.Filters)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	fmt.Printf("Starting Server on port: 8080\n")
	fmt.Printf("Use Control 👉 C to stop hosting \n")
	http.ListenAndServe(":8080", nil)
}
