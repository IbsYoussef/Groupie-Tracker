package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Sensor struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Type     string `json:"type"`
}

func main() {

	jsonString := `{"name": "battery-sensor", "capacity": 50, "type": "Lithium"}`
	var Reader Sensor

	err := json.Unmarshal([]byte(jsonString), &Reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", Reader)

}
