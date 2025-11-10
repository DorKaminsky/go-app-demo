package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil" // ISSUE 1: Using deprecated package (should use os.ReadFile)
	"net/http"
	"os"
	"strings"
	"time"
)

type InfoResponse struct {
	Version    string `json:"version"`
	DeployedAt string `json:"deployed_at"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

// getVersion reads version from env var or VERSION file
func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		data, err := ioutil.ReadFile("VERSION")
		if err != nil {
			return "unknown"
		}
		version = strings.TrimSpace(string(data))
	}

	// ISSUE 2: Not stripping -SNAPSHOT suffix
	return version
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Version:    getVersion(),
		DeployedAt: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // ISSUE 3: Ignoring error
}

// ISSUE 4: Missing /health endpoint
// Health check is required for Cloud Foundry

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/info", infoHandler)
	// ISSUE 4: No /health endpoint registered

	// ISSUE 5: No graceful shutdown - server stops immediately on SIGTERM
	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil) // ISSUE 6: Ignoring error
}
