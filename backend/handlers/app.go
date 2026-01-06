package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<span class="status-value">✅ Running</span>`)
}

func GetTrackNumberFromBothPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nbTracksDeezer := getTrackNumberFromDeezer(w)
	nbTracksSpotify := getTrackNumberFromSpotify(w)

	result := map[string]interface{}{
		"status":				"success",
		"nb_tracks_Deezer":   	nbTracksDeezer.Total,
		"nb_tracks_Spotify":    nbTracksSpotify.Tracks.Total,
	}

	json.NewEncoder(w).Encode(result)
	return
}

func GetSyncStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")



	// Simulate sync status retrieval
	result := map[string]interface{}{
		"status":       "success",
		"sync_status":  "All playlists are synchronised",
	}

	json.NewEncoder(w).Encode(result)
}