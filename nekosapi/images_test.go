package nekosapi

import (
	"slices"
	"testing"
)

// Test GetImages()
func TestGetImages(t *testing.T) {
	imageParams := GetImagesParams{
		Ratings: []Rating{SAFE, SUGGESTIVE},
		Tags:    []string{"blue_hair", "skirt"},
		Limit:   20,
	}

	images, err := GetImages(imageParams)
	if err != nil {
		t.Fatal(err)
	}

	for _, image := range images.Items {
		if image.Rating != SAFE && image.Rating != SUGGESTIVE {
			t.Fatal("Expecting only SAFE or SUGGESTIVE images")
		}
		if !slices.Contains(image.Tags, "blue_hair") {
			t.Fatal("Expecting blue_hair tag to be included in tags")
		}
		if !slices.Contains(image.Tags, "skirt") {
			t.Fatal("Expecting skirt tag to be included in tags")
		}
	}
}

// Test GetImageById()
func TestGetImageById(t *testing.T) {
	id := 1449

	const expectedTagName string = "purple_hair"

	result, err := GetImageById(id)
	if err != nil {
		t.Fatal(err)
	}

	containsPurpleHair := false
	for _, tag := range result.Tags {
		if tag == expectedTagName {
			containsPurpleHair = true
		}
	}

	if !containsPurpleHair {
		t.Fatal("purple_hair tag expected to be in this image")
	}
}

// Test GetImageFileById()
func TestGetImageFileById(t *testing.T) {
	id := 1449

	imageFile, err := GetImageFileById(id)
	if err != nil {
		t.Fatal(err)
	}

	if imageFile == "" {
		t.Fatal("Image file was not found")
	}
}

// Test GetRandomImages()
func TestGetRandomImages(t *testing.T) {
	imageParams := GetRandomImagesParams{
		Ratings: []Rating{SAFE, SUGGESTIVE},
		Limit:   3,
	}

	images, err := GetRandomImages(imageParams)
	if err != nil {
		t.Fatal(err)
	}

	for _, image := range images {
		if image.Rating != SAFE && image.Rating != SUGGESTIVE {
			t.Fatal("Expecting only SAFE or SUGGESTIVE images")
		}
	}

	if len(images) != imageParams.Limit {
		t.Fatal("Number of returned images does not match limit")
	}
}

// Test GetRandomFile()
func TestGetRandomFile(t *testing.T) {
	randomImageParams := GetRandomFileParams{
		Ratings: []Rating{SAFE},
	}

	randomImageFile, err := GetRandomFile(randomImageParams)
	if err != nil {
		t.Fatal(err)
	}

	if randomImageFile == "" {
		t.Fatal("Random file was not found")
	}
}
