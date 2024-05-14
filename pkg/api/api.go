package api

import(
	"encoding/json"
    "net/http"
)

type Payload struct {
	Status string
	IsRetrograde bool `json:"is_retrograde"`
}

func FetchMercuryRetrogradeStatus() (bool, error) {
    url := "https://mercuryretrogradeapi.com"

    response, error := http.Get(url)
    if error != nil {
        return false, error
    }
    defer response.Body.Close()

    var result Payload
    if error := json.NewDecoder(response.Body).Decode(&result); error != nil {
        return false, error
    }

    return result.IsRetrograde, nil
}
