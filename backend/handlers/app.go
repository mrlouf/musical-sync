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

	close(spotify)
	close(deezer)

	var sync_status string
	if nbTracksDeezer.Total == nbTracksSpotify.Tracks.Total {
		sync_status = "Playlists are synchronised"
	} else {
		sync_status = "Playlists are not synchronised"
	}

	result := map[string]interface{}{
		"status":       "success",
		"sync_status":  sync_status,
	}

	json.NewEncoder(w).Encode(result)
}