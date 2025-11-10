package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type InfoResponse struct {
	Version    string `json:"version"`
	DeployedAt string `json:"deployed_at"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

// getVersion reads version from environment variable or VERSION file
// Normalizes version by removing -SNAPSHOT suffix for production
func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		data, err := os.ReadFile("VERSION")
		if err != nil {
			log.Printf("Warning: Could not read VERSION file: %v", err)
			return "unknown"
		}
		version = strings.TrimSpace(string(data))
	}

	// Normalize version by removing -SNAPSHOT suffix
	version = strings.TrimSuffix(version, "-SNAPSHOT")
	return version
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Version:    getVersion(),
		DeployedAt: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Served /info endpoint - version: %s", response.Version)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set up routes
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/health", healthHandler)

	// Create server with timeouts
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down gracefully...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}