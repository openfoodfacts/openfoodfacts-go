package openfoodfacts

import (
	"testing"
)



func TestQuerySearch_PrintsResults(t *testing.T) {
	api := NewClient("world", "", "")

	query := SearchQuery{
		Keyword:      "Lindt Excellence 70",
		CategoryTags: []string{"en:dark-chocolates"},
		Page:         1,
		PageSize:     5,
	}

	result, err := api.QuerySearch(query)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Count: %d", result.Count)
	t.Logf("Page: %d", result.Page)
	t.Logf("Hits: %d", len(result.Hits))
	for _, hit := range result.Hits {
		t.Logf("- %s | %s", hit["code"], hit["product_name"])
	}
}

func TestClient_newRequest_UserAgent(t *testing.T) {
	api := NewClient("world", "", "")
	api.Sandbox()

	if got := api.newRequest("GET", "/api/v0/product/%s.json", "5201051001076"); got.Header.Get("User-Agent") != defaultUserAgent {
		t.Errorf("newRequest() = %v, want %v", got.Header.Get("User-Agent"), defaultUserAgent)
	}

	customUserAgent := "CoolFoodApp - Android - Version 1.0 - https://coolfoodapp.com"
	api.UserAgent(customUserAgent)

	if got := api.newRequest("GET", "/api/v0/product/%s.json", "5201051001076"); got.Header.Get("User-Agent") != customUserAgent {
		t.Errorf("newRequest() = %v, want %v", got.Header.Get("User-Agent"), customUserAgent)
	}
}
