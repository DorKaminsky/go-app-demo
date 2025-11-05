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

	// Verify -SNAPSHOT is stripped
	expectedVersion := "1.189.0"
	if response.Version != expectedVersion {
		t.Errorf("version should be %s (with -SNAPSHOT stripped), got %s", expectedVersion, response.Version)
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}
}

func TestInfoHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler should reject POST: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
	}{
		{
			name:     "strips SNAPSHOT from env var",
			envValue: "2.0.0-SNAPSHOT",
			expected: "2.0.0",
		},
		{
			name:     "handles version without SNAPSHOT",
			envValue: "3.0.0",
			expected: "3.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("VERSION", tt.envValue)
			defer os.Unsetenv("VERSION")

			version := getVersion()
			if version != tt.expected {
				t.Errorf("getVersion() = %v, want %v", version, tt.expected)
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "UP" {
		t.Errorf("health status should be UP, got %s", response.Status)
	}
}

func TestHealthHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler should reject POST: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
