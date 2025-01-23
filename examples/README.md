# Examples

## Get images for a character

```go
package main

import (
	"fmt"

	"github.com/generalvamp/go-nekos-api/nekosapi"
)

func main() {
	imageParams := nekosapi.GetImagesParams{
		Ratings: []nekosapi.Rating{nekosapi.SAFE, nekosapi.SUGGESTIVE},
		Tags:    []string{"black_hair", "dress"},
		Limit:   2,
		Offset:  1,
	}

	images, err := nekosapi.GetImages(imageParams)
	if err != nil {
		fmt.Println(err)
	}

	for _, image := range images.Items {
		fmt.Printf("Image: %v\n", image.URL)
	}
}

// Output:
// Image: https://s3.nyeki.dev/nekos-api/images/original/110eb935-e399-4872-bcc1-a3db4fcbd651.webp
// Image: https://s3.nyeki.dev/nekos-api/images/original/77f8b866-359e-46ed-80fc-59ce54787d36.webp
```

| 1st Image   | 2nd Image |
| ---------   | ----------|
| ![image 1](https://s3.nyeki.dev/nekos-api/images/original/110eb935-e399-4872-bcc1-a3db4fcbd651.webp)| ![image 2](https://s3.nyeki.dev/nekos-api/images/original/77f8b866-359e-46ed-80fc-59ce54787d36.webp)  |

## Get random images

```go
package main

import (
	"fmt"

	"github.com/generalvamp/go-nekos-api/nekosapi"
)

func main() {
	randomImageParams := nekosapi.GetRandomImagesParams{
		Ratings: []nekosapi.Rating{nekosapi.SAFE, nekosapi.SUGGESTIVE},
		Limit:   3,
	}

	images, err := nekosapi.GetRandomImages(randomImageParams)
	if err != nil {
		fmt.Println(err)
	}

	for _, image := range images {
		fmt.Printf("Image: %v\n", image.URL)
	}
}

// Output:
// Image: https://s3.nyeki.dev/nekos-api/images/original/d45b7242-fcd7-444b-9e18-86767de05463.webp
// Image: https://s3.nyeki.dev/nekos-api/images/original/c2d0bc3b-ead0-4ebb-9c78-c9c05b975bb0.webp
// Image: https://s3.nyeki.dev/nekos-api/images/original/ad824fdd-1df4-473b-be1b-49d7ffb9b3a3.webp
```

| 1st Image   | 2nd Image   | 3nd Image |
| ---------   | ---------   | --------- |
| ![image 1](https://s3.nyeki.dev/nekos-api/images/original/d45b7242-fcd7-444b-9e18-86767de05463.webp)    | ![image 2](https://s3.nyeki.dev/nekos-api/images/original/c2d0bc3b-ead0-4ebb-9c78-c9c05b975bb0.webp)| ![image 3](https://s3.nyeki.dev/nekos-api/images/original/ad824fdd-1df4-473b-be1b-49d7ffb9b3a3.webp) |

