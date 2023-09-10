package main

import (
	"fmt"
	"log"
	"net/http"
)

type API struct {
	ID        int
	Artists   Artists
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func main() {

	Artists_API, Aerr := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if Aerr != nil {
		log.Print(Aerr)
	}
	defer Artists_API.Body.Close()

	Locations_API, Lerr := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if Lerr != nil {
		log.Print(Lerr)
	}
	defer Locations_API.Body.Close()

	Dates_API, Derr := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if Derr != nil {
		log.Print(Derr)
	}
	defer Dates_API.Body.Close()

	Relations_API, Rerr := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if Rerr != nil {
		log.Print(Rerr)
	}
	defer Relations_API.Body.Close()

	fileserver := http.FileServer(http.Dir("."))
	http.Handle("/", fileserver)

	fmt.Printf("Starting Server on port: 8080\n")
	fmt.Printf("Use Control 👉 C to stop hosting \n")
	http.ListenAndServe(":8080", nil)

}
