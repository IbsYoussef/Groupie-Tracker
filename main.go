package main

import (
	"encoding/json"
	"log"
)

// type Artists struct {
// 	ID           int      `json:"id"`
// 	Image        string   `json:"image"`
// 	Name         string   `json:"name"`
// 	Members      []string `json:"members"`
// 	CreationDate int      `json:"creationDate"`
// 	FirstAlbum   string   `json:"firstAlbum"`
// 	Locations    string   `json:"locations"`
// 	ConcertDates string   `json:"concertDates"`
// 	Relations    string   `json:"relations"`
// }

type Jsonreader struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     int    `json:"time"`
}

func main() {

	// API_endpoint, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// if err != nil {
	// 	log.Print(err)
	// }

	jsonstring := `{"name": "battery-sensor", "capacity": "50", "time": "22"}`

	var reader Jsonreader
	err := json.Unmarshal([]byte(jsonstring), &reader)
	if err != nil {
		log.Fatal(err)
	}

}
