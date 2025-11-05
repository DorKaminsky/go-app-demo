package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInfoHandler(t *testing.T) {
	// Set version for testing
	os.Setenv("VERSION", "1.189.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

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
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// ISSUE: This test will fail because version should be "1.189.0" not "1.189.0-SNAPSHOT"
	// The test expects the bug to exist
	if response.Version == "" {
		t.Error("version should not be empty")
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}
}

func TestGetVersion(t *testing.T) {
	// Test with environment variable
	os.Setenv("VERSION", "2.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	version := getVersion()
	if version == "" {
		t.Error("version should not be empty")
	}

	// ISSUE: Should test that -SNAPSHOT is stripped, but doesn't
}
