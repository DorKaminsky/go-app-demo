package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestInfoHandler(t *testing.T) {
	// Test with normalized version (without -SNAPSHOT)
	os.Setenv("VERSION", "1.189.0")
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

	if response.Version == "" {
		t.Error("version should not be empty")
	}

	if response.Version != "1.189.0" {
		t.Errorf("expected version 1.189.0, got %s", response.Version)
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}

	// Verify Content-Type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
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

	if response.Status != "healthy" {
		t.Errorf("expected status 'healthy', got '%s'", response.Status)
	}

	// Verify Content-Type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}
}

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     string
	}{
		{
			name:     "clean version from env",
			envValue: "2.0.0",
			want:     "2.0.0",
		},
		{
			name:     "version with -SNAPSHOT suffix gets normalized",
			envValue: "2.0.0-SNAPSHOT",
			want:     "2.0.0",
		},
		{
			name:     "another snapshot version",
			envValue: "1.189.0-SNAPSHOT",
			want:     "1.189.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("VERSION", tt.envValue)
			defer os.Unsetenv("VERSION")

			version := getVersion()
			if version != tt.want {
				t.Errorf("getVersion() = %v, want %v", version, tt.want)
			}
		})
	}
}

func TestGetVersionFromFile(t *testing.T) {
	// Unset VERSION env to force reading from file
	os.Unsetenv("VERSION")

	// Create temporary VERSION file
	content := []byte("1.189.0\n")
	err := os.WriteFile("VERSION_TEST", content, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("VERSION_TEST")

	// Temporarily rename real VERSION file if it exists
	versionExists := false
	if _, err := os.Stat("VERSION"); err == nil {
		versionExists = true
		os.Rename("VERSION", "VERSION.bak")
		defer os.Rename("VERSION.bak", "VERSION")
	}

	// Create VERSION file for test
	err = os.WriteFile("VERSION", []byte("1.189.0\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	if !versionExists {
		defer os.Remove("VERSION")
	}

	version := getVersion()
	if version != "1.189.0" {
		t.Errorf("expected version from file to be 1.189.0, got %s", version)
	}
}

func TestGetVersionFileNotFound(t *testing.T) {
	// Unset VERSION env
	os.Unsetenv("VERSION")

	// Temporarily rename VERSION file if it exists
	versionExists := false
	if _, err := os.Stat("VERSION"); err == nil {
		versionExists = true
		err := os.Rename("VERSION", "VERSION.bak")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Rename("VERSION.bak", "VERSION")
	}

	version := getVersion()
	if version != "unknown" {
		t.Errorf("expected 'unknown' when VERSION file doesn't exist, got %s", version)
	}
}

func TestVersionNormalization(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "snapshot version is normalized",
			input: "1.189.0-SNAPSHOT",
			want:  "1.189.0",
		},
		{
			name:  "clean version stays the same",
			input: "1.189.0",
			want:  "1.189.0",
		},
		{
			name:  "multiple snapshots removed",
			input: "2.0.0-SNAPSHOT",
			want:  "2.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSuffix(tt.input, "-SNAPSHOT")
			if result != tt.want {
				t.Errorf("normalization failed: got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestInfoHandlerReturnsNormalizedVersion(t *testing.T) {
	// Set version with -SNAPSHOT
	os.Setenv("VERSION", "3.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)

	var response InfoResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Verify -SNAPSHOT was removed
	if strings.Contains(response.Version, "-SNAPSHOT") {
		t.Errorf("version should not contain -SNAPSHOT, got %s", response.Version)
	}

	if response.Version != "3.0.0" {
		t.Errorf("expected normalized version 3.0.0, got %s", response.Version)
	}
}