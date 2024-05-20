package api

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Status string
	IMIR bool `json:"is_retrograde"`
}

func Fetch() (bool, error) {
    url := "https://mercuryretrogradeapi.com"

    response, error := http.Get(url)
    // TODO: Handle errors.
    if error != nil {
        return false, error
    }
    defer response.Body.Close()

    var result Payload
    if error := json.NewDecoder(response.Body).Decode(&result); error != nil {
        return false, error
    }

    return result.IMIR, nil
}

func FetchByDateInput(givenDate string) (bool, error) {
    url := "https://mercuryretrogradeapi.com?date=" + givenDate

    response, error := http.Get(url)
    if error != nil {
        return false, error
    }
    defer response.Body.Close()

    var result Payload
    if error := json.NewDecoder(response.Body).Decode(&result); error != nil {
        return false, error
    }

    return result.IMIR, nil
}
