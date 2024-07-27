package nekosapi

import (
	"testing"
)

// Test GetCharacters()
func TestGetCharacters(t *testing.T) {
	const expectedCharacter string = "Tohsaka, Rin"

	characterParams := GetCharactersParams{
		Search: "Tohsa",
		Limit:  5,
	}

	characters, err := GetCharacters(characterParams)
	if err != nil {
		t.Fatal(err)
	}

	containsExpectedCharacter := false
	for _, artist := range characters.Items {
		if artist.Name == expectedCharacter {
			containsExpectedCharacter = true
		}
	}

	if !containsExpectedCharacter {
		t.Fatalf("Expected character %v not found when searching character", expectedCharacter)
	}
}

// Test GetCharacterById()
func TestGetCharacterById(t *testing.T) {
	id := 19
	const expectedCharacter string = "Kitagawa, Marin"

	result, err := GetCharacterById(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.Name != expectedCharacter {
		t.Fatalf("Expected character %v, but found character %v", expectedCharacter, result.Name)
	}
}

// Test GetCharacterImages()
func TestGetCharacterImages(t *testing.T) {
	id := 20
	const expectedCharacter string = "Lumine"

	images, err := GetCharacterImages(id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, image := range images.Items {
		containsExpectedCharacter := false
		for _, character := range image.Characters {
			if character.Name == expectedCharacter {
				containsExpectedCharacter = true
			}
		}

		if !containsExpectedCharacter {
			t.Fatalf("Not all images contain the expected character %v", expectedCharacter)
		}
	}
}
