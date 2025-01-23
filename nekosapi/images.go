package nekosapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Rating represents the (age) rating of an image.
type Rating string

// Valid Ratings
const (
	SAFE       Rating = "safe"
	SUGGESTIVE Rating = "suggestive"
	BORDERLINE Rating = "borderline"
	EXPLICIT   Rating = "explicit"
)

// URL for the /images endpoint
const IMAGES_ENDPOINT string = BASE_URL_V4 + "/images"

// Image represents an image
type Image struct {
	ID            int      `json:"id"`
	URL           string   `json:"url"`
	Rating        Rating   `json:"rating"`
	ColorDominant []int    `json:"color_dominant"`
	ColorPalette  [][]int  `json:"color_palette"`
	ArtistName    string   `json:"artist_name"`
	Tags          []string `json:"tags"`
	SourceURL     string   `json:"source_url"`
}

// PaginatedImage represents paginated Image results
type PaginatedImage struct {
	Items []Image `json:"items"`
	Count int     `json:"count"`
}

// GetImagesParams represents params used by GetImages()
type GetImagesParams struct {
	Ratings []Rating
	Artist  string
	Tags    []string
	Limit   int
	Offset  int
}

// GetRandomImagesParams represents params used by GetRandomImages()
type GetRandomImagesParams struct {
	Ratings []Rating
	Artist  string
	Tags    []string
	Limit   int
}

// GetRandomFileParams represents params used by GetRandomFile()
type GetRandomFileParams struct {
	Tags    []string
	Ratings []Rating
	Artist  string
}

// GetImages() corresponds to the /images endpoint.
//
// This endpoint allows you to search for an image, filtering by tags, characters, artists, etc.
func GetImages(params GetImagesParams) (*PaginatedImage, error) {
	endpointURL := IMAGES_ENDPOINT

	values := url.Values{}

	if len(params.Ratings) > 0 {
		var ratings []string
		for _, r := range params.Ratings {
			ratings = append(ratings, string(r))
		}

		values.Add("rating", strings.Join(ratings, ","))
	}

	if params.Artist != "" {
		values.Add("artist_name", params.Artist)
	}

	if len(params.Tags) > 0 {
		values.Add("tags", strings.Join(params.Tags, ","))
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	if params.Offset < 0 {
		return nil, errors.New("param offset must be 0 or higher")
	}

	values.Add("offset", strconv.Itoa(params.Offset))

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedImages := &PaginatedImage{}
	err := getRequest(urlWithParams, paginatedImages)
	if err != nil {
		return nil, err
	}

	return paginatedImages, nil
}

// GetImageById() corresponds to the /images/{id} endpoint.
//
// This endpoint allows you to get an image by its ID.
func GetImageById(id int) (*Image, error) {
	endpointURL := IMAGES_ENDPOINT

	finalUrl := fmt.Sprintf("%v/%d", endpointURL, id)

	image := &Image{}
	err := getRequest(finalUrl, image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetImageFileById() corresponds to the /images/{image_id}/file endpoint.
//
// This endpoint allows you to get a redirect to the image's file URL, based on the given image id.
func GetImageFileById(id int) (string, error) {
	endpointURL := IMAGES_ENDPOINT

	finalUrl := fmt.Sprintf("%v/%d%v", endpointURL, id, FILE_PATH)

	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return "", err
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)

	if response != nil && response.StatusCode == http.StatusFound {
		url, err := response.Location()
		if err != nil {
			return "", err
		}

		return url.String(), nil
	}

	return "", err
}

// GetRandomImages() corresponds to the /images/random endpoint.
//
// This endpoint allows you to get x random images, filtering by tags, characters, artists, etc.
func GetRandomImages(params GetRandomImagesParams) (*PaginatedImage, error) {
	endpointURL := IMAGES_ENDPOINT + RANDOM_PATH

	values := url.Values{}

	if len(params.Ratings) > 0 {
		var ratings []string
		for _, r := range params.Ratings {
			ratings = append(ratings, string(r))
		}

		values.Add("rating", strings.Join(ratings, ","))
	}

	if params.Artist != "" {
		values.Add("artist_name", params.Artist)
	}

	if len(params.Tags) > 0 {
		values.Add("tags", strings.Join(params.Tags, ","))
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedImages := &PaginatedImage{}
	err := getRequest(urlWithParams, paginatedImages)
	if err != nil {
		return nil, err
	}

	return paginatedImages, nil
}

// GetRandomFile() corresponds to the /images/random/file endpoint.
//
// This endpoint allows you to get a redirect to a random image's file URL, filtering by tags, characters, artists, etc.
func GetRandomFile(params GetRandomFileParams) (string, error) {
	endpointURL := IMAGES_ENDPOINT + RANDOM_PATH + FILE_PATH

	values := url.Values{}

	if len(params.Ratings) > 0 {
		var ratings []string
		for _, r := range params.Ratings {
			ratings = append(ratings, string(r))
		}

		values.Add("rating", strings.Join(ratings, ","))
	}

	if params.Artist != "" {
		values.Add("artist_name", params.Artist)
	}

	if len(params.Tags) > 0 {
		values.Add("tags", strings.Join(params.Tags, ","))
	}

	urlWithParams := endpointURL + "?" + values.Encode()

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return "", err
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)

	if response != nil && response.StatusCode == http.StatusFound {
		url, err := response.Location()
		if err != nil {
			return "", err
		}

		return url.String(), nil
	}

	if response != nil && response.StatusCode != http.StatusFound {
		return "", errors.New(fmt.Sprintf("Status Code:%d,Status Text:%v", response.StatusCode, http.StatusText(response.StatusCode)))
	}

	return "", nil
}
