package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInfoHandler(t *testing.T) {
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

	if response.Version == "" {
		t.Error("version should not be empty")
	}

	if response.DeployedAt == "" {
		t.Error("deployed_at should not be empty")
	}
}

func TestGetVersion(t *testing.T) {
	os.Setenv("VERSION", "2.0.0-SNAPSHOT")
	defer os.Unsetenv("VERSION")

	version := getVersion()
	if version == "" {
		t.Error("version should not be empty")
	}
}

func TestGetVersionFromFile(t *testing.T) {
	os.Unsetenv("VERSION")
	f, err := os.Create("VERSION")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("VERSION")
	_, err = f.WriteString("3.0.0-PROD\n")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	version := getVersion()
	if version != "3.0.0-PROD" {
		t.Errorf("expected version '3.0.0-PROD', got '%s'", version)
	}
}

func TestGetVersionFileMissing(t *testing.T) {
	os.Unsetenv("VERSION")
	os.Remove("VERSION") // ensure file is missing
	version := getVersion()
	if version != "unknown" {
		t.Errorf("expected version 'unknown', got '%s'", version)
	}
}

func TestInfoHandlerMissingVersionFile(t *testing.T) {
	os.Unsetenv("VERSION")
	os.Remove("VERSION")
	req, _ := http.NewRequest("GET", "/info", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
	var response InfoResponse
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.Version != "unknown" {
		t.Errorf("expected version 'unknown', got '%s'", response.Version)
	}
}

func TestInfoHandlerContentType(t *testing.T) {
	os.Setenv("VERSION", "1.0.0")
	req, _ := http.NewRequest("GET", "/info", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", ct)
	}
}
