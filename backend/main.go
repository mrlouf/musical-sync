package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	// Start background polling
	go startPolling()

	// Set up routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/sync-status", syncStatusHandler)
	http.HandleFunc("/poll/deezer", pollDeezerHandler)
	http.HandleFunc("/poll/spotify", pollSpotifyHandler)

	log.Printf("Backend server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// healthHandler returns the health status of the backend
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<span class="status-value">✅ Running</span>`)
}

// syncStatusHandler returns the current synchronization status
func syncStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	lastSyncText := "Never"
	if !syncStatus.LastSync.IsZero() {
		lastSyncText = syncStatus.LastSync.Format("2006-01-02 15:04:05")
	}

	html := fmt.Sprintf(`
		<div class="sync-status">
			<h2 style="margin-bottom: 20px; color: #333;">Sync Status</h2>
			<div class="status-item">
				<span class="status-label">Backend Status:</span>
				<span class="status-value">%s</span>
			</div>
			<div class="status-item">
				<span class="status-label">Deezer Connection:</span>
				<span class="status-value">%s</span>
			</div>
			<div class="status-item">
				<span class="status-label">Spotify Connection:</span>
				<span class="status-value">%s</span>
			</div>
			<div class="status-item">
				<span class="status-label">Last Sync:</span>
				<span class="status-value">%s</span>
			</div>
		</div>
	`, syncStatus.BackendStatus, syncStatus.DeezerStatus, syncStatus.SpotifyStatus, lastSyncText)

	fmt.Fprint(w, html)
}

// pollDeezerHandler simulates polling Deezer API
func pollDeezerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate API call
	result := map[string]interface{}{
		"status":  "success",
		"message": "Deezer API polled successfully",
		"data": map[string]interface{}{
			"playlists":    []string{"Favorites", "Road Trip", "Workout"},
			"last_updated": time.Now().Format(time.RFC3339),
		},
	}

	syncStatus.DeezerStatus = "✅ Connected"
	syncStatus.LastSync = time.Now()

	json.NewEncoder(w).Encode(result)
}

// pollSpotifyHandler simulates polling Spotify API
func pollSpotifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate API call
	result := map[string]interface{}{
		"status":  "success",
		"message": "Spotify API polled successfully",
		"data": map[string]interface{}{
			"playlists":    []string{"Discover Weekly", "Daily Mix", "Liked Songs"},
			"last_updated": time.Now().Format(time.RFC3339),
		},
	}

	syncStatus.SpotifyStatus = "✅ Connected"
	syncStatus.LastSync = time.Now()

	json.NewEncoder(w).Encode(result)
}

// startPolling begins background polling of external APIs
func startPolling() {
	syncStatus.IsPolling = true
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	log.Println("Starting background polling service...")

	for range ticker.C {
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
