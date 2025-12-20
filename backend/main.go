package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mod/handlers/handlers"
)

// SyncStatus represents the synchronization status across platforms
type SyncStatus struct {
	BackendStatus string    `json:"backend_status"`
	DeezerStatus  string    `json:"deezer_status"`
	SpotifyStatus string    `json:"spotify_status"`
	LastSync      time.Time `json:"last_sync"`
	IsPolling     bool      `json:"is_polling"`
}

var (
	syncStatus = SyncStatus{
		BackendStatus: "Running",
		DeezerStatus:  "Not connected",
		SpotifyStatus: "Not connected",
		LastSync:      time.Time{},
		IsPolling:     false,
	}
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start background polling
	go startPolling(ctx)

	// Set up routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/sync-status", syncStatusHandler)
	http.HandleFunc("/poll/deezer", pollDeezerHandler)
	http.HandleFunc("/poll/spotify", pollSpotifyHandler)
	http.HandleFunc("/login/deezer", loginDeezerHandler)
	http.HandleFunc("/login/spotify", loginSpotifyHandler)

	// Create server
	server := &http.Server{
		Addr: ":" + port,
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Backend server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Cancel the polling context
	cancel()

	// Gracefully shutdown the server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}



// startPolling begins background polling of external APIs
func startPolling(ctx context.Context) {
	syncStatus.IsPolling = true
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	log.Println("Starting background polling service...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping background polling service...")
			syncStatus.IsPolling = false
			return
		case <-ticker.C:
			log.Println("Polling external APIs...")

			// In a real implementation, this would call actual APIs
			// For now, we just log the polling activity

			if syncStatus.DeezerStatus == "✅ Connected" {
				log.Println("Checking Deezer sync status...")
			}

			if syncStatus.SpotifyStatus == "✅ Connected" {
				log.Println("Checking Spotify sync status...")
			}
		}
	}
}
