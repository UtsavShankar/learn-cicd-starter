package auth

import (
	"net/http"
	"testing"
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
