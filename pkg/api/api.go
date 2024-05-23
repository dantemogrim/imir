package api

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Status string
	IMIR   bool `json:"is_retrograde"`
}

func Fetch() (bool, error) {
	url := "https://mercuryretrogradeapi.com"

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var result Payload
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.IMIR, nil
}

func FetchByDateInput(givenDate string) (bool, error) {
	url := "https://mercuryretrogradeapi.com?date=" + givenDate

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var result Payload
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.IMIR, nil
}
