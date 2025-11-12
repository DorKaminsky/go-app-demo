package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Test version normalization
	os.Setenv("VERSION", "1.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	version := getVersion()

	// Should strip -SNAPSHOT
	if version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", version)
	}
}

func TestGetVersionFromFile(t *testing.T) {
	// Clear env var to test file reading
	os.Unsetenv("VERSION")

	// Create temporary VERSION file
	content := []byte("2.0.0-SNAPSHOT\n")
	err := os.WriteFile("VERSION", content, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("VERSION")

	version := getVersion()

	// This test will fail because it expects wrong value
	if version != "3.0.0" {
		t.Errorf("Expected version '3.0.0', got '%s'", version)
	}
}

func TestInfoHandler(t *testing.T) {
	os.Setenv("VERSION", "1.0.0")
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
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", response.Version)
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
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
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}
