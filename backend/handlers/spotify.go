package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"log"

	"backend/utils"
	"backend/models"
)

// TODO : Implement actual Spotify OAuth flow
func LoginSpotifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate login process
	result := map[string]interface{}{
		"status":  "success",
		"message": "Spotify login simulated successfully",
	}

	json.NewEncoder(w).Encode(result)
}


func getTrackNumberFromSpotify(w http.ResponseWriter) models.SpotifyTrackNumberResponse {

	spotifyPlaylistID := os.Getenv("SPOTIFY_PLAYLIST_ID")
	if spotifyPlaylistID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "SPOTIFY_PLAYLIST_ID not set in environment",
		})
		return models.SpotifyTrackNumberResponse{}
	}
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", spotifyPlaylistID)

	req, err := http.NewRequest("GET", url, nil)
	token := utils.GenerateSpotifyToken()
	fmt.Println("Generated Spotify Token:", token)
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch playlist from Spotify",
		})
		return models.SpotifyTrackNumberResponse{}
	}
	defer resp.Body.Close()

	var nbTracksSpotify models.SpotifyTrackNumberResponse
	json.NewDecoder(resp.Body).Decode(&nbTracksSpotify)

	return nbTracksSpotify
}

func GetSpotifyPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	spotifyPlaylistID := os.Getenv("SPOTIFY_PLAYLIST_ID")
	if spotifyPlaylistID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "SPOTIFY_PLAYLIST_ID not set in environment",
		})
		return
	}
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", spotifyPlaylistID)

	req, err := http.NewRequest("GET", url, nil)
	token := utils.GenerateSpotifyToken()
	fmt.Println("Generated Spotify Token:", token)
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Failed to fetch playlist from Spotify",
		})
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Spotify API response status: %s\n", resp.Status)

	var playlistData struct {
		Items []struct {
			Track struct {
				Name   string `json:"name"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
				Album struct {
					Name string `json:"name"`
				} `json:"album"`
			} `json:"track"`
		} `json:"items"`
	}

	json.NewDecoder(resp.Body).Decode(&playlistData)

	fmt.Printf("%v\n", playlistData)
	fmt.Printf("Fetched %d tracks from Spotify playlist\n", len(playlistData.Items))
	// fmt.Printf("%v\n", playlistData.Items[0].Track.Name) // Print the name of the first track for debugging

	result := map[string]interface{}{
		"status":	"success",
		"tracks":	playlistData.Items,
	}

	json.NewEncoder(w).Encode(result)
}