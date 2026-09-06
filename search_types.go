package openfoodfacts

type SearchResponse struct {
	// Success fields
	Hits         []map[string]interface{} `json:"hits,omitempty"`
	Aggregations map[string]interface{}   `json:"aggregations,omitempty"`
	Facets       map[string]FacetInfo     `json:"facets,omitempty"`
	Charts       map[string]interface{}   `json:"charts,omitempty"`
	Page         int                      `json:"page,omitempty"`
	PageSize     int                      `json:"page_size,omitempty"`
	PageCount    int                      `json:"page_count,omitempty"`
	Took         int                      `json:"took,omitempty"`
	TimedOut     bool                     `json:"timed_out,omitempty"`
	Count        int                      `json:"count,omitempty"`
	IsCountExact bool                     `json:"is_count_exact,omitempty"`
	Warnings     []SearchResponseError    `json:"warnings,omitempty"`

	// Error fields
	Errors []SearchResponseError `json:"errors,omitempty"`

	// Shared
	Debug SearchResponseDebug `json:"debug"`
}

type FacetInfo struct {
	Name             string      `json:"name"`
	Items            []FacetItem `json:"items"`
	CountErrorMargin *int        `json:"count_error_margin,omitempty"`
}

type FacetItem struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Count    *int   `json:"count"`
	Selected bool   `json:"selected"`
}

type SearchResponseDebug struct {
	Query map[string]interface{} `json:"query"`
}

type SearchResponseError struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}