package nekosapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Base URL for version 3 of Nekos API
const BASE_URL_V3 string = "https://api.nekosapi.com/v3"

// get_request() takes in a url and parses the json into parsedData
func get_request(url string, parsedData any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &parsedData)

	return nil
}
