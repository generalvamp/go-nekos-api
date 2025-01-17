// Package nekosapi is a go wrapper for using [Nekos API].
//
// [Nekos API]: https://nekosapi.com/
package nekosapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Base URL for version 4 of Nekos API.
const BASE_URL_V4 string = "https://api.nekosapi.com/v4"

// Paths used for building URLs
const (
	IMAGES_PATH string = "/images"
	RANDOM_PATH string = "/random"
	FILE_PATH   string = "/file"
)

// Constants for use with limit query parameter
const (
	MIN_LIMIT = 1   // Minimum value able to be used with limit paramter
	MAX_LIMIT = 100 // Maximum value able to be used with limit paramter
)

// getRequest() takes in a url and parses the json into parsedData.
func getRequest(url string, parsedData any) error {
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
