package nekosapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// Character represents character that may be present in an image
type Character struct {
	ID          int      `json:"id"`
	IDV2        string   `json:"id_v2"`
	Name        string   `json:"name"`
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Ages        []int    `json:"ages"`
	Height      *int     `json:"height"`
	Weight      *int     `json:"weight"`
	Gender      string   `json:"gender"`
	Species     string   `json:"species"`
	Birthday    string   `json:"birthday"`
	Nationality string   `json:"nationality"`
	Occupations []string `json:"occupations"`
}

// PaginatedArtist represents paginated Artist results
type PaginatedCharacter struct {
	Items []Character `json:"items"`
	Count int         `json:"count"`
}

// GetCharactersParams represents params used by GetCharacters()
type GetCharactersParams struct {
	Search      string
	Age         *int
	Gender      string
	Species     string
	Nationality string
	Occupation  string
	Limit       int
	Offset      int
}

const CHARACTERS_ENDPOINT string = BASE_URL_V3 + "/characters"

// GetCharacters() corresponds to /characters endpoint.
// This endpoint allows you to search for a character, filtering by name, aliases, description, etc.
func GetCharacters(params GetCharactersParams) (*PaginatedCharacter, error) {
	endpointURL := CHARACTERS_ENDPOINT

	values := url.Values{}

	if params.Search != "" {
		values.Add("search", params.Search)
	}

	if params.Age != nil {
		values.Add("age", strconv.Itoa(*params.Age))
	}

	if params.Gender != "" {
		values.Add("gender", params.Gender)
	}

	if params.Species != "" {
		values.Add("species", params.Species)
	}

	if params.Nationality != "" {
		values.Add("nationality", params.Nationality)
	}

	if params.Occupation != "" {
		values.Add("occupation", params.Occupation)
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	values.Add("offset", strconv.Itoa(params.Offset))

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedCharacter := &PaginatedCharacter{}
	err := get_request(urlWithParams, paginatedCharacter)
	if err != nil {
		return nil, err
	}

	return paginatedCharacter, nil
}

// GetCharacterById() corresponds to /characters/{id} endpoint.
// This endpoint allows you to get an character by its ID.
func GetCharacterById(id int) (*Character, error) {
	endpointURL := CHARACTERS_ENDPOINT

	finalUrl := fmt.Sprintf("%v/%d", endpointURL, id)

	character := &Character{}
	err := get_request(finalUrl, character)
	if err != nil {
		return nil, err
	}

	return character, nil
}

// GetCharacterImages() corresponds to the /characters/{id}/images endpoint.
// This endpoint allows you to get all images picturing a certain character.
func GetCharacterImages(id int, limit int, offset int) (*PaginatedImage, error) {
	endpointURL := CHARACTERS_ENDPOINT

	values := url.Values{}

	if limit != 0 {
		if limit < MIN_LIMIT || limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(limit))
	}

	values.Add("offset", strconv.Itoa(offset))

	finalUrl := fmt.Sprintf("%v/%d%v", endpointURL, id, IMAGES_URL)

	paginatedImage := &PaginatedImage{}
	err := get_request(finalUrl, paginatedImage)
	if err != nil {
		return nil, err
	}

	return paginatedImage, nil
}
