package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func GetJSON(url string, target interface{}) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API JSON responded with Error Code %d", response.StatusCode)
	}

	return json.NewDecoder(response.Body).Decode(target)
}

func AppendToStruct() {
	for i := range locations.Index {
		API[i].Locations.Index = append(API[i].Locations.Index, locations.Index[i])
	}
	for i := range dates.Index {
		API[i].Dates.Index = append(API[i].Dates.Index, dates.Index[i])
	}
	for i := range relations.Index {
		API[i].Relations.Index = append(API[i].Relations.Index, relations.Index[i])
	}
}

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	Dates        Dates
	Relations    Relations
}

type Locations struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}
type Date struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Index []Relation `json:"index"`
}
type Relation struct {
	ID             int64                  `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

var tpl *template.Template
var client *http.Client
var API []Artist
var locations Locations
var dates Dates
var relations Relations

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	GetJSON("https://groupietrackers.herokuapp.com/api/artists", &API)
	GetJSON("https://groupietrackers.herokuapp.com/api/locations", &locations)
	GetJSON("https://groupietrackers.herokuapp.com/api/dates", &dates)
	GetJSON("https://groupietrackers.herokuapp.com/api/relation", &relations)
	AppendToStruct()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/web", webHandler)
	http.HandleFunc("/card", cardHandler)

	assets := http.FileServer(http.Dir("Assets"))
	http.Handle("/Assets/", http.StripPrefix("/Assets/", assets))

	fmt.Printf("Listening... on port 👉 :8080 \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error code: %s", err.Error())
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404, page not found", 404)
	} else if r.Method != "GET" {
		http.Error(w, "Error 500, server endpoint not located", 500)
	} else {
		tpl.ExecuteTemplate(w, "index.html", API)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "Error 404, page not found", 404)
	} else if r.Method != "GET" {
		http.Error(w, "Error 500, server endpoint not located", 500)
	} else {
		tpl.ExecuteTemplate(w, "about.html", nil)
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/web" {
		http.Error(w, "Error 404, page not found", 404)
	} else if r.Method != "GET" {
		http.Error(w, "Error 500, server endpoint not located", 500)
	} else {
		tpl.ExecuteTemplate(w, "web.html", API)
	}
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/card" {
		http.Error(w, "Error 404, page not found", 404)
	} else if r.Method != "GET" {
		http.Error(w, "Error 500, server endpoint not located", 500)
	} else {
		tpl.ExecuteTemplate(w, "card-hover.html", API)
	}
}
