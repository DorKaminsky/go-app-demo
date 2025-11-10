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

// InfoResponse represents the version and deployment time returned by the /info endpoint.
type InfoResponse struct {
	Version    string `json:"version"`
	DeployedAt string `json:"deployed_at"`
}

func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		data, err := ioutil.ReadFile("VERSION")
		if err != nil {
			return "unknown"
		}
		version = strings.TrimSpace(string(data))
	}
	// Normalize: strip -SNAPSHOT or any suffix after a dash
	if idx := strings.Index(version, "-"); idx != -1 {
		version = version[:idx]
	}
	return version
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Version:    getVersion(),
		DeployedAt: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
