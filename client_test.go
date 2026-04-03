package openfoodfacts

import (
	"testing"
)



func TestQuerySearch_PrintsResults(t *testing.T) {
	api := NewClient("world", "", "")

	query := SearchQuery{
		Lang:         "en",
		CategoryTags: []string{"beverages"},
		Page:         1,
		PageSize:     2,
	}

	result, err := api.QuerySearch(query)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Count: %d", result.Count)
	t.Logf("Hits returned: %d", len(result.Hits))
	for _, hit := range result.Hits {
		cats, _ := hit["categories_tags"].([]interface{})
		t.Logf("- %s | categories: %v", hit["product_name"], cats)
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
