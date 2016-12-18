package openfoodfacts

type ProductResult struct {
	StatusVerbose string   `json:"status_verbose"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	Product       *Product `json:"product"`
}
