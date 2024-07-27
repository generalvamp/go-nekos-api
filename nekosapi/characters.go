package nekosapi

// Character represents character that may be present in an image
type Character struct {
	ID          int      `json:"id"`
	IDV2        string   `json:"id_v2"`
	Name        string   `json:"name"`
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Ages        []int    `json:"ages"`
	Height      *int     `json:"height"`
	Weight      *int     `json:"weight"`
	Gender      string   `json:"gender"`
	Species     string   `json:"species"`
	Birthday    string   `json:"birthday"`
	Nationality string   `json:"nationality"`
	Occupations []string `json:"occupations"`
}
