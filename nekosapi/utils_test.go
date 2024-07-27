package nekosapi

import "testing"

// Test getRequest()
func TestGetRequest(t *testing.T) {
	image := &Image{}
	err := getRequest(IMAGES_ENDPOINT+"/1", image)
	if err != nil {
		t.Error(err)
	}

	if image.ImageURL == "" {
		t.Error("Get request failed to parse image")
	}
}
