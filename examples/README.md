# Examples

## Get images for a character

```go
package main

import (
	"fmt"

	"github.com/generalvamp/go-nekos-api/nekosapi"
)

func main() {
	params := nekosapi.GetCharactersParams{
		Search: "Kurumi",
	}

	// Get characters based on search param Kurumi
	paginatedCharacters, err := nekosapi.GetCharacters(params)
	if err != nil {
		panic(err)
	}

	// loop over the found characters
	for _, character := range paginatedCharacters.Items {

		// Get two images for the current character
		paginatedImages, err := nekosapi.GetCharacterImages(character.ID, 2, 0)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Character Name: %v\n", character.Name)

		for _, image := range paginatedImages.Items {
			fmt.Printf("Image: %v\n", image.ImageURL)
		}
	}
}

// Output:
// Character Name: Toksaki, Kurumi
// Image: https://cdn.nekosapi.com/images/original/452f652f-704a-409e-9ec0-2874684a4152.webp
// Image: https://cdn.nekosapi.com/images/original/71b0dd4e-f660-49f9-ae6f-cddab2ec0eed.webp
```

| Character Name    | 1st Image   | 2nd Image |
| ---------------   | ---------   | ----------|
| Toksaki, Kurumi     | ![image 1](https://cdn.nekosapi.com/images/original/452f652f-704a-409e-9ec0-2874684a4152.webp)| ![image 2](https://cdn.nekosapi.com/images/original/71b0dd4e-f660-49f9-ae6f-cddab2ec0eed.webp) |

## Get random images

```go
package main

import (
	"fmt"

	"github.com/generalvamp/go-nekos-api/nekosapi"
)

func main() {
	isFlagged := false
	params := nekosapi.GetRandomImagesParams{
		Ratings:   []nekosapi.Rating{nekosapi.SAFE},
		IsFlagged: &isFlagged,
		Limit:     3,
	}

	// Get random images based on params
	paginatedImages, err := nekosapi.GetRandomImages(params)
	if err != nil {
		panic(err)
	}

	// loop over the found images
	for _, image := range paginatedImages.Items {
		fmt.Printf("Image: %v\n", image.ImageURL)
	}
}
// Output:
// Image: https://cdn.nekosapi.com/images/original/0947f334-940f-40bf-a857-166a3f38dcca.webp
// Image: https://cdn.nekosapi.com/images/original/c2dfbc28-1cad-4f94-8ff2-501b4b60ff74.webp
// Image: https://cdn.nekosapi.com/images/original/129fc4b2-8beb-458b-bdfd-5af0b7818247.webp
```

| 1st Image   | 2nd Image   | 3nd Image |
| ---------   | ---------   | --------- |
| ![image 1](https://cdn.nekosapi.com/images/original/0947f334-940f-40bf-a857-166a3f38dcca.webp)    | ![image 2](https://cdn.nekosapi.com/images/original/c2dfbc28-1cad-4f94-8ff2-501b4b60ff74.webp)| ![image 3](https://cdn.nekosapi.com/images/original/129fc4b2-8beb-458b-bdfd-5af0b7818247.webp) |

## Get images for a tag
```go
package main

import (
	"fmt"

	"github.com/generalvamp/go-nekos-api/nekosapi"
)

func main() {
	// Get tag based on search sword
	paginatedTags, err := nekosapi.GetTags("sword", nil, 1, 0)
	if err != nil {
		panic(err)
	}

	swordTag := paginatedTags.Items[0]

	isFlagged := false
	params := nekosapi.GetImagesParams{
		Ratings:   []nekosapi.Rating{nekosapi.SAFE},
		IsFlagged: &isFlagged,
		Tag:       []int{swordTag.ID},
		Limit:     2,
	}

	// Get images based on params
	paginatedImages, err := nekosapi.GetImages(params)
	if err != nil {
		panic(err)
	}

	// loop over the found images
	for _, image := range paginatedImages.Items {
		fmt.Printf("Image: %v\n", image.ImageURL)
	}
}
// Output:
// Image: https://cdn.nekosapi.com/images/original/ae5fcccd-bb35-4fa3-b6e8-12ac3b7ab10a.webp
// Image: https://cdn.nekosapi.com/images/original/ddf0998c-7422-4f10-9ab7-536bc0be3d2d.webp
```

|  1st Image   | 2nd Image |
| -----------  | --------- |
| ![image 1](https://cdn.nekosapi.com/images/original/ae5fcccd-bb35-4fa3-b6e8-12ac3b7ab10a.webp)| ![image 2](https://cdn.nekosapi.com/images/original/ddf0998c-7422-4f10-9ab7-536bc0be3d2d.webp) |