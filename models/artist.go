package models

type Artist struct {
	ID           int64     `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int64     `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    Locations `json:"-"`
	Dates        Dates     `json:"-"`
	Relations    Relations `json:"-"`
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

var API []Artist
var LocationsData Locations
var DatesData Dates
var RelationsData Relations

func AppendToStruct() {
	for i := 0; i < len(API); i++ {
		if i < len(LocationsData.Index) {
			API[i].Locations.Index = append(API[i].Locations.Index, LocationsData.Index[i])
		}
		if i < len(DatesData.Index) {
			API[i].Dates.Index = append(API[i].Dates.Index, DatesData.Index[i])
		}
		if i < len(RelationsData.Index) {
			API[i].Relations.Index = append(API[i].Relations.Index, RelationsData.Index[i])
		}
	}
}
