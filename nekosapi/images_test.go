package nekosapi

// import "testing"

// // Test GetImages()
// func TestGetImages(t *testing.T) {
// 	imageParams := GetImagesParams{
// 		Ratings: []Rating{SAFE},
// 		Limit:   5,
// 	}

// 	images, err := GetImages(imageParams)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, image := range images.Items {
// 		if image.Rating != SAFE {
// 			t.Fatal("Expecting only SAFE images")
// 		}
// 	}
// }

// // Test GetImageById()
// func TestGetImageById(t *testing.T) {
// 	id := 1449

// 	const expectedCharacterName string = "Tohsaka, Rin"
// 	const expectedTagName string = "Purple hair"

// 	result, err := GetImageById(id)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	containsRin := false
// 	for _, character := range result.Characters {
// 		if character.Name == expectedCharacterName {
// 			containsRin = true
// 		}
// 	}

// 	containsPurpleHair := false
// 	for _, tag := range result.Tags {
// 		if tag.Name == expectedTagName {
// 			containsPurpleHair = true
// 		}
// 	}

// 	if !containsRin {
// 		t.Fatal("Rin expected to be in this image")
// 	}

// 	if !containsPurpleHair {
// 		t.Fatal("Purple hair tag expected to be in this image")
// 	}
// }

// // Test GetRandomImages()
// func TestGetRandomImages(t *testing.T) {
// 	isOriginal := true

// 	imageParams := GetRandomImagesParams{
// 		Ratings:    []Rating{SAFE},
// 		Limit:      5,
// 		IsOriginal: &isOriginal,
// 	}

// 	images, err := GetRandomImages(imageParams)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, image := range images.Items {
// 		if image.Rating != SAFE {
// 			t.Fatal("Expecting only SAFE images")
// 		}
// 		if !image.IsOriginal {
// 			t.Fatal("Expecting only original images")
// 		}
// 	}
// }

// // Test GetImageArtist()
// func TestGetImageArtist(t *testing.T) {
// 	const expectedArtist string = "Rosuuri"

// 	artist, err := GetImageArtist(13494)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if artist.Name != expectedArtist {
// 		t.Fatalf("Expected artist %v", expectedArtist)
// 	}
// }

// // Test GetTags()
// func TestGetTags(t *testing.T) {
// 	isNSFW := false

// 	tags, err := GetTags("", &isNSFW, 10, 10)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, tag := range tags.Items {
// 		if tag.IsNSFW {
// 			t.Fatalf("Expecting only SFW tags. Found tag %v", tag.Name)
// 		}
// 	}
// }

// // Test GetTagById()
// func TestGetTagById(t *testing.T) {
// 	const expectedTag string = "Beach"
// 	const expectedTagId int = 20

// 	tag, err := GetTagById(expectedTagId)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if tag.Name != expectedTag {
// 		t.Fatalf("Expecting tag %v, found tag %v, %v", expectedTag, tag.Name, tag)
// 	}
// }

// // Test GetTagImages()
// func TestGetTagImages(t *testing.T) {
// 	const expectedTag string = "Beach"
// 	const expectedTagId int = 20

// 	images, err := GetTagImages(expectedTagId, 5, 0)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, image := range images.Items {
// 		containsExpectedTag := false
// 		for _, tag := range image.Tags {
// 			if tag.Name == expectedTag {
// 				containsExpectedTag = true
// 			}
// 		}

// 		if !containsExpectedTag {
// 			t.Fatalf("All images do not contain the expected tag %v", expectedTag)
// 		}
// 	}
// }
