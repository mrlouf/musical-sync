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

    var allTracks []models.SpotifyPlaylistItem
    limit := 100
    offset := 0

    token := utils.GenerateSpotifyToken()
    client := &http.Client{}

    for {
        url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?limit=%d&offset=%d", spotifyPlaylistID, limit, offset)
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            log.Fatal(err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  "error",
                "message": "Failed to create request",
            })
            return
        }
        req.Header.Set("Authorization", "Bearer "+token)

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

        var playlistData models.SpotifyPlaylistResponse
        if err := json.NewDecoder(resp.Body).Decode(&playlistData); err != nil {
            log.Fatal(err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  "error",
                "message": "Failed to decode Spotify response",
            })
            return
        }

        allTracks = append(allTracks, playlistData.Items...)

        if len(playlistData.Items) < limit {
			fmt.Println("All tracks fetched from Spotify playlist.")
            break
        }
        offset += limit
    }

    result := map[string]interface{}{
        "status": "success",
        "tracks": allTracks,
    }

    json.NewEncoder(w).Encode(result)
}