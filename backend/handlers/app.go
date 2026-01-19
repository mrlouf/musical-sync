package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"backend/models"
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

	spotify := make (chan models.SpotifyTrackNumberResponse)
	deezer := make (chan models.DeezerTrackNumberResponse)

	go func() {
		spotify <- getTrackNumberFromSpotify(w)
	}()

	go func() {
		deezer <- getTrackNumberFromDeezer(w)
	}()

	nbTracksSpotify := <- spotify
	nbTracksDeezer := <- deezer

	fmt.Println("Spotify Tracks:", nbTracksSpotify.Tracks.Total)
	fmt.Println("Deezer Tracks:", nbTracksDeezer.Total)

	if nbTracksSpotify.Tracks.Total != nbTracksDeezer.Total {
		result := map[string]interface{}{
			"status":       "success",
			"sync_status":  "Playlists are not synchronised",
		}

		close(spotify)
		close(deezer)

		json.NewEncoder(w).Encode(result)
		return
	}

	close(spotify)
	close(deezer)

	result := map[string]interface{}{
		"status":       "success",
		"sync_status":  "All playlists are synchronised",
	}

	json.NewEncoder(w).Encode(result)
}