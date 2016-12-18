package openfoodfacts

type Ingredient struct {
	ID   string `json:"id"`
	Rank int    `json:"rank,omitempty"`
	Text string `json:"text"`
}
