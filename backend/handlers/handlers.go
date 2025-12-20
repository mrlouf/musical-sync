package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

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

func loginDeezerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate login process
	result := map[string]interface{}{
		"status":  "success",
		"message": "Deezer login simulated successfully",
	}

	json.NewEncoder(w).Encode(result)
}

func loginSpotifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate login process
	result := map[string]interface{}{
		"status":  "success",
		"message": "Spotify login simulated successfully",
	}

	json.NewEncoder(w).Encode(result)
}