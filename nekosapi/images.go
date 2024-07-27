package nekosapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
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

// GetImagesParams represents params used by GetImages()
type GetImagesParams struct {
	Ratings      []Rating
	IsOriginal   *bool
	IsScreenshot *bool
	IsFlagged    *bool
	IsAnimated   *bool
	Artist       *int
	Character    []int
	Tag          []int
	Limit        int
	Offset       int
}

// GetRandomImagesParams represents params used by GetRandomImages()
type GetRandomImagesParams struct {
	Ratings      []Rating
	IsOriginal   *bool
	IsScreenshot *bool
	IsFlagged    *bool
	IsAnimated   *bool
	Artist       []int
	Character    []int
	Tag          []int
	Limit        int
}

// PaginatedImage represents paginated Image results
type PaginatedImage struct {
	Items []Image `json:"items"`
	Count int     `json:"count"`
}

// PaginatedTag represents paginated Tag results
type PaginatedTag struct {
	Items []Tag `json:"items"`
	Count int   `json:"count"`
}

// Image represents an image
type Image struct {
	ID             int         `json:"id"`
	IDV2           string      `json:"id_v2"`
	ImageURL       string      `json:"image_url"`
	SampleURL      string      `json:"sample_url"`
	ImageSize      int         `json:"image_size"`
	ImageWidth     int         `json:"image_width"`
	ImageHeight    int         `json:"image_height"`
	SampleSize     int         `json:"sample_size"`
	SampleWidth    int         `json:"sample_width"`
	SampleHeight   int         `json:"sample_height"`
	Source         string      `json:"source"`
	SourceID       *int        `json:"source_id"`
	Rating         Rating      `json:"rating"`
	Verification   string      `json:"verification"`
	HashMd5        string      `json:"hash_md5"`
	HashPerceptual string      `json:"hash_perceptual"`
	ColorDominant  []int       `json:"color_dominant"`
	ColorPalette   [][]int     `json:"color_palette"`
	Duration       *int        `json:"duration"`
	IsOriginal     bool        `json:"is_original"`
	IsScreenshot   bool        `json:"is_screenshot"`
	IsFlagged      bool        `json:"is_flagged"`
	IsAnimated     bool        `json:"is_animated"`
	Artist         Artist      `json:"artist"`
	Characters     []Character `json:"characters"`
	Tags           []Tag       `json:"tags"`
	CreatedAt      float64     `json:"created_at"`
	UpdatedAt      float64     `json:"updated_at"`
}

// Tag represents a tag for categorizing images
type Tag struct {
	ID          int    `json:"id"`
	IDV2        string `json:"id_v2"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sub         string `json:"sub"`
	IsNSFW      bool   `json:"is_nsfw"`
}

// TODO: maybe rename these and or group them in one const ()
const IMAGES_ENDPOINT string = BASE_URL_V3 + "/images"

const IMAGES_URL string = "/images"
const RANDOM_URL string = "/random"
const FILE_URL string = "/file"
const ARTIST_URL string = "/artist"
const CHARACTERS_URL string = "/characters"
const TAGS_URL string = "/tags"
const REPORT_URL string = "/report"

const MIN_LIMIT = 1
const MAX_LIMIT = 100

// GetImages() corresponds to /images endpoint.
// This endpoint allows you to search for an image, filtering by tags, characters, artists, etc.
func GetImages(params GetImagesParams) (*PaginatedImage, error) {
	endpointURL := IMAGES_ENDPOINT

	values := url.Values{}

	for _, r := range params.Ratings {
		values.Add("rating", string(r))
	}

	if params.IsOriginal != nil {
		values.Add("is_original", strconv.FormatBool(*params.IsOriginal))
	}

	if params.IsScreenshot != nil {
		values.Add("is_screenshot", strconv.FormatBool(*params.IsScreenshot))
	}

	if params.IsFlagged != nil {
		values.Add("is_flagged", strconv.FormatBool(*params.IsFlagged))
	}

	if params.IsAnimated != nil {
		values.Add("is_animated", strconv.FormatBool(*params.IsAnimated))
	}

	if params.Artist != nil {
		values.Add("artist", strconv.Itoa(*params.Artist))
	}

	for _, c := range params.Character {
		values.Add("character", strconv.Itoa(c))
	}

	for _, t := range params.Tag {
		values.Add("tag", strconv.Itoa(t))
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	values.Add("offset", strconv.Itoa(params.Offset))

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedImages := &PaginatedImage{}
	err := get_request(urlWithParams, paginatedImages)
	if err != nil {
		return nil, err
	}

	return paginatedImages, nil
}

// GetImageById() corresponds to /images/{id} endpoint.
// This endpoint allows you to get an image by its ID.
func GetImageById(id int) (*Image, error) {
	endpointURL := IMAGES_ENDPOINT

	finalUrl := fmt.Sprintf("%v/%d", endpointURL, id)

	image := &Image{}
	err := get_request(finalUrl, image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetRandomImages() corresponds to the images/random endpoint.
// This endpoint allows you to get x random images, filtering by tags, characters, artists, etc.
func GetRandomImages(params GetRandomImagesParams) (*PaginatedImage, error) {
	endpointURL := IMAGES_ENDPOINT + RANDOM_URL

	values := url.Values{}

	for _, r := range params.Ratings {
		values.Add("rating", string(r))
	}

	if params.IsOriginal != nil {
		values.Add("is_original", strconv.FormatBool(*params.IsOriginal))
	}

	if params.IsScreenshot != nil {
		values.Add("is_screenshot", strconv.FormatBool(*params.IsScreenshot))
	}

	if params.IsFlagged != nil {
		values.Add("is_flagged", strconv.FormatBool(*params.IsFlagged))
	}

	if params.IsAnimated != nil {
		values.Add("is_animated", strconv.FormatBool(*params.IsAnimated))
	}

	for _, r := range params.Artist {
		values.Add("artist", strconv.Itoa(r))
	}

	for _, c := range params.Character {
		values.Add("character", strconv.Itoa(c))
	}

	for _, t := range params.Tag {
		values.Add("tag", strconv.Itoa(t))
	}

	if params.Limit != 0 {
		if params.Limit < MIN_LIMIT || params.Limit > MAX_LIMIT {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(params.Limit))
	}

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedImages := &PaginatedImage{}
	err := get_request(urlWithParams, paginatedImages)
	if err != nil {
		return nil, err
	}

	return paginatedImages, nil
}

// GetImageArtist() corresponds to the /images/{id}/artist endpoint.
// This endpoint allows you to get an image's artist.
func GetImageArtist(id int) (*Artist, error) {
	finalUrl := fmt.Sprintf("%v/%d%v", IMAGES_ENDPOINT, id, ARTIST_URL)

	artist := &Artist{}
	err := get_request(finalUrl, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

// GetTags() corresponds to /images/tags endpoint
// This endpoint allows you to search for a tag, filtering by name, description, and whether it's NSFW or not.
func GetTags(search string, isNSFW *bool, limit int, offset int) (*PaginatedTag, error) {
	endpointURL := IMAGES_ENDPOINT + TAGS_URL

	values := url.Values{}

	if search != "" {
		values.Add("search", search)
	}

	if isNSFW != nil {
		values.Add("is_nsfw", strconv.FormatBool(*isNSFW))
	}

	if limit != 0 {
		if limit < 1 || limit > 100 {
			return nil, errors.New("param limit must be between 1 and 100")
		}

		values.Add("limit", strconv.Itoa(limit))
	}

	values.Add("offset", strconv.Itoa(offset))

	urlWithParams := endpointURL + "?" + values.Encode()

	paginatedTag := &PaginatedTag{}
	err := get_request(urlWithParams, paginatedTag)
	if err != nil {
		return nil, err
	}

	return paginatedTag, nil
}

// GetTagById() corresponds to endpoint /images/tags/{id}
// This endpoint allows you to get a tag by its ID.
func GetTagById(id int) (*Tag, error) {
	endpointURL := IMAGES_ENDPOINT + TAGS_URL

	finalUrl := fmt.Sprintf("%v/%d", endpointURL, id)

	tag := &Tag{}
	err := get_request(finalUrl, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

// GetTagImages() corresponds to the /images/tags/{id}/images endpoint.
// This endpoint allows you to search for a tag, filtering by name, description, and whether it's NSFW or not.
func GetTagImages(id int, limit int, offset int) (*PaginatedImage, error) {
	endpointURL := IMAGES_ENDPOINT + TAGS_URL

	values := url.Values{}

	if limit != 0 {
		if limit < 1 || limit > 100 {
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
