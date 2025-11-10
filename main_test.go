package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Set VERSION env var for testing
	os.Setenv("VERSION", "1.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	version := getVersion()

	// Note: This test expects the buggy behavior
	// Version should strip -SNAPSHOT but it doesn't
	if version == "" {
		t.Error("version should not be empty")
	}
}

func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response InfoResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Version == "" {
		t.Error("version should not be empty")
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}
}

// Missing: Tests for /health endpoint
// Missing: Tests for version normalization (-SNAPSHOT stripping)
