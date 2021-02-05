package openfoodfacts

import (
	"testing"
)

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
