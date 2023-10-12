package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Artist struct {
	Name             string `json:"name"`
	Age              int    `json:"age"`
	Date_of_Creation int    `json:"date"`
	Location_ID      int    `json:"id"`
	Tour             Tour   `json:"tour"`
}

type Tour struct {
	Location string  `json:"location"`
	Time     string  `json:"time"`
	Date     string  `json:"date"`
	Price    float32 `json:"price"`
}

func main() {

	tour := Tour{Location: "London", Time: "18:00", Date: "20th July", Price: 35.55}
	artist := Artist{Name: "Freddy Mercury", Age: 30, Date_of_Creation: 1980, Location_ID: 6540, Tour: tour}

	byteArray, err := json.MarshalIndent(artist, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteArray))

}
