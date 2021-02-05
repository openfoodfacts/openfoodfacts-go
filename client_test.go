package openfoodfacts

import (
	"testing"
)

func TestClient_newRequest_UserAgent(t *testing.T) {
	userAgent := "CoolFoodApp - Android - Version 1.0 - https://coolfoodapp.com"

	api := NewClient("world", "", "")
	api.Sandbox()

	if got := api.newRequest("GET", "/api/v0/product/%s.json", "5201051001076"); got.Header.Get("User-Agent") != "" {
		t.Errorf("newRequest() = %v, want %v", got.Header.Get("User-Agent"), "")
	}

	api.UserAgent(userAgent)

	if got := api.newRequest("GET", "/api/v0/product/%s.json", "5201051001076"); got.Header.Get("User-Agent") != userAgent {
		t.Errorf("newRequest() = %v, want %v", got.Header.Get("User-Agent"), userAgent)
	}
}
