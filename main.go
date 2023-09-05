package main

import (
	"log"
	"net/http"
)

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {

	API_endpoint, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Print(err)
	}

}
