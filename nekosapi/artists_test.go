package nekosapi

import (
	"testing"
)

// Test GetArtists()
func TestGetArtists(t *testing.T) {
	const expectedArtist string = "Rosuuri"

	artistsParams := GetArtistsParams{
		Search: "Rosu",
		Limit:  5,
	}

	artists, err := GetArtists(artistsParams)
	if err != nil {
		t.Fatal(err)
	}

	containsExpectedArtist := false
	for _, artist := range artists.Items {
		if artist.Name == expectedArtist {
			containsExpectedArtist = true
		}
	}

	if !containsExpectedArtist {
		t.Fatalf("Expected artist %v not found when searching artists", expectedArtist)
	}
}

// Test GetArtistById()
func TestGetArtistById(t *testing.T) {
	id := 9
	const expectedArtist string = "Rosuuri"

	result, err := GetArtistById(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.Name != expectedArtist {
		t.Fatalf("Expected artist %v, but found artist %v", expectedArtist, result.Name)
	}
}

// Test GetArtistImages()
func TestGetArtistImages(t *testing.T) {
	id := 9
	const expectedArtist string = "Rosuuri"

	images, err := GetArtistImages(id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, image := range images.Items {
		if image.Artist.Name != expectedArtist {
			t.Fatalf("All images do not contain the expected tag %v", expectedArtist)
		}
	}
}
