package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// ISSUE 1: Using deprecated ioutil package (should use os.ReadFile)
// ISSUE 2: Missing error handling in multiple places
// ISSUE 3: No logging middleware or structured logging

type InfoResponse struct {
	Version    string `json:"version"`
	DeployedAt string `json:"deployed_at"`
}

// ISSUE 4: Version parsing bug - doesn't strip -SNAPSHOT suffix
func getVersion() string {
	// Try environment variable first
	version := os.Getenv("VERSION")
	if version != "" {
		return version
	}

	// Read from VERSION file
	data, err := ioutil.ReadFile("VERSION")
	if err != nil {
		return "unknown"
	}

	// ISSUE 5: Should strip -SNAPSHOT but doesn't
	return strings.TrimSpace(string(data))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// ISSUE 6: No method check - should only allow GET
	
	response := InfoResponse{
		Version:    getVersion(),
		DeployedAt: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // ISSUE 7: Ignoring error from Encode
}

// ISSUE 8: No health check endpoint for Cloud Foundry
// ISSUE 9: No graceful shutdown handling

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/info", infoHandler)

	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil) // ISSUE 10: Ignoring error return
}
