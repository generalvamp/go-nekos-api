package nekosapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// URL for the /artists endpoint
const ARTISTS_ENDPOINT string = BASE_URL_V3 + "/artists"

// Artist represents the artist of an image
type Artist struct {
	ID           int      `json:"id"`
	IDV2         string   `json:"id_v2"`
	Name         string   `json:"name"`
	Aliases      []string `json:"aliases"`
	ImageURL     string   `json:"image_url"`
	Links        []string `json:"links"`
	PolicyRepost *bool    `json:"policy_repost"`
	PolicyCredit bool     `json:"policy_credit"`
	PolicyAi     bool     `json:"policy_ai"`
}

// PaginatedArtist represents paginated Artist results
type PaginatedArtist struct {
	Items []Artist `json:"items"`
	Count int      `json:"count"`
}

// GetArtistsParams represents params used by GetArtists()
type GetArtistsParams struct {
	Search       string
	PolicyRepost *bool
	PolicyCredit *bool
	PolicyAI     *bool
	Limit        int
	Offset       int
}

// GetArtists() correspond to the /artists endpoint.
//
// This endpoint allows you to search for an artist, filtering by name, aliases, policies, etc.
func GetArtists(params GetArtistsParams) (*PaginatedArtist, error) {
	endpointURL := ARTISTS_ENDPOINT

	values := url.Values{}

	if params.Search != "" {
		values.Add("search", params.Search)
	}

	if params.PolicyRepost != nil {
		values.Add("policy_repost", strconv.FormatBool(*params.PolicyRepost))
	}

	if params.PolicyCredit != nil {
		values.Add("policy_credit", strconv.FormatBool(*params.PolicyCredit))
	}

	if params.PolicyRepost != nil {
		values.Add("policy_repost", strconv.FormatBool(*params.PolicyRepost))
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	values.Add("offset", strconv.Itoa(params.Offset))

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedArtist := &PaginatedArtist{}
	err := getRequest(urlWithParams, paginatedArtist)
	if err != nil {
		return nil, err
	}

	return paginatedArtist, nil
}

// GetArtistById() corresponds to the /artists/{id} endpoint.
//
// This endpoint allows you to get an artist by its ID.
func GetArtistById(id int) (*Artist, error) {
	endpointURL := ARTISTS_ENDPOINT

	finalUrl := fmt.Sprintf("%v/%d", endpointURL, id)

	artist := &Artist{}
	err := getRequest(finalUrl, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

// GetArtistImages() corresponds to the /artists/{id}/images endpoint.
//
// This endpoint allows you to get all images made by an artist.
func GetArtistImages(id int, limit int, offset int) (*PaginatedImage, error) {
	endpointURL := ARTISTS_ENDPOINT

	values := url.Values{}

	if limit != 0 {
		if limit < MIN_LIMIT || limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(limit))
	}

	values.Add("offset", strconv.Itoa(offset))

	finalUrl := fmt.Sprintf("%v/%d%v", endpointURL, id, IMAGES_PATH)

	paginatedImage := &PaginatedImage{}
	err := getRequest(finalUrl, paginatedImage)
	if err != nil {
		return nil, err
	}

	return paginatedImage, nil
}
