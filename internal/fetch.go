package internal

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

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

func LoadData() error {
	if err := GetJSON("https://groupietrackers.herokuapp.com/api/artists", &models.API); err != nil {
		return err
	}

	if err := GetJSON("https://groupietrackers.herokuapp.com/api/locations", &models.LocationsData); err != nil {
		return err
	}
	if err := GetJSON("https://groupietrackers.herokuapp.com/api/dates", &models.DatesData); err != nil {
		return err
	}
	if err := GetJSON("https://groupietrackers.herokuapp.com/api/relation", &models.RelationsData); err != nil {
		return err
	}

	models.AppendToStruct()
	return nil
}
