package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"log"

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

func getTrackNumberFromDeezer(w http.ResponseWriter) models.DeezerTrackNumberResponse {
	deezerPlaylistID := os.Getenv("DEEZER_PLAYLIST_ID")
	if deezerPlaylistID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "DEEZER_PLAYLIST_ID not set in environment",
		})
		return models.DeezerTrackNumberResponse{}
	}

	url := fmt.Sprintf("https://api.deezer.com/playlist/%s", deezerPlaylistID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch playlist from Deezer",
		})
		return models.DeezerTrackNumberResponse{}
	}
	defer resp.Body.Close()

	var nbTracksDeezer models.DeezerTrackNumberResponse
	json.NewDecoder(resp.Body).Decode(&nbTracksDeezer)

	return nbTracksDeezer
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

func GetRandomDeezerTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://api.deezer.com/track/3208591241" // Example track ID

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch track from Deezer",
		})
		return
	}
	defer resp.Body.Close()

	var trackResponse models.DeezerTrackResponse
	json.NewDecoder(resp.Body).Decode(&trackResponse)

	result := map[string]interface{}{
		"status":	"success",
		"track":	trackResponse.Title,
		"artist":	trackResponse.Artist.Name,
		"album":	trackResponse.Album.Title,
	}

	json.NewEncoder(w).Encode(result)
}

func GetRandomDeezerAlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	url := "https://api.deezer.com/album/705023831" // Example album ID

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch album from Deezer",
		})
		return
	}
	defer resp.Body.Close()

	var albumResponse models.DeezerAlbumResponse
	json.NewDecoder(resp.Body).Decode(&albumResponse)

	result := map[string]interface{}{
		"status":	"success",
		"album":	albumResponse.Title,
		"artist":	albumResponse.Artist.Name,
		"tracks":	albumResponse.Tracks.Data,
	}

	json.NewEncoder(w).Encode(result)
}

func GetDeezerPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deezerPlaylistID := os.Getenv("DEEZER_PLAYLIST_ID")
	if deezerPlaylistID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "DEEZER_PLAYLIST_ID not set in environment",
		})
		return
	}

	url := fmt.Sprintf("https://api.deezer.com/playlist/%s/tracks?limit=1000", deezerPlaylistID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch playlist tracks from Deezer",
		})
		return
	}
	defer resp.Body.Close()

	var tracklistResponse models.DeezerTracklistResponse
	json.NewDecoder(resp.Body).Decode(&tracklistResponse)

	result := map[string]interface{}{
		"status":	"success",
		"title":	tracklistResponse.Title,
		"tracks":	tracklistResponse.Data,
	}

	json.NewEncoder(w).Encode(result)
}