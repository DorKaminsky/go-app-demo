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

type InfoResponse struct {
	Version    string `json:"version"`
	DeployedAt string `json:"deployed_at"`
}

func getVersion() string {
	version := os.Getenv("VERSION")
	if version != "" {
		return version
	}

	data, err := ioutil.ReadFile("VERSION")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(data))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Version:    getVersion(),
		DeployedAt: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/info", infoHandler)

	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
