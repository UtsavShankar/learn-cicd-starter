package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")
	result, err := GetAPIKey(headers)
	if result != "my-secret-key" || err != nil {
		t.Errorf("failed %s", result)
	}
}

func TestGetAPIKeyNoAuth(t *testing.T) {
	headers := http.Header{}
	headers.Set("something else", "ApiKey my-secret-key")
	result, err := GetAPIKey(headers)
	if result == "my-secret-key" || err == nil {
		t.Errorf("failed %s", result)
	}
}
