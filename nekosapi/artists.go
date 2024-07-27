package nekosapi

// Artist represents the artist of an image
type Artist struct {
	ID           int      `json:"id"`
	IDV2         string   `json:"id_v2"`
	Name         string   `json:"name"`
	Aliases      []string `json:"aliases"`
	ImageURL     string   `json:"image_url"`
	Links        []string `json:"links"`
	PolicyRepost *bool    `json:"policy_repost"`
	PolicyCredit bool     `json:"policy_credit"`
	PolicyAi     bool     `json:"policy_ai"`
}
