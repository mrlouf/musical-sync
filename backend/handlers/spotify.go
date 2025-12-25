package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"

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
	token := utils.GenerateSpotifyToken()
	client := &http.Client{}
	var playlistData []models.SpotifyPlaylistResponse

    for url != "" {
        fmt.Printf("Fetching Spotify playlist from URL: %s\n", url)
        
        req, _ := http.NewRequest("GET", url, nil)
        // TODO handle error
        req.Header.Set("Authorization", "Bearer " + token)

        resp, err := client.Do(req)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  "error",
                "message": "Failed to fetch playlist from Spotify",
            })
            return
        }

        fmt.Printf("Spotify API response status: %s\n", resp.Status)

        var pageData struct {
            Items []models.SpotifyPlaylistResponse `json:"items"`
            Next  string                           `json:"next"`
        }

        // read full body to inspect actual JSON
        body, err := io.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            log.Println("read body error:", err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status": "error",
                "message": "Failed to read Spotify response",
            })
            return
        }
        fmt.Printf("Spotify response body: %s\n", string(body))

        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
            log.Println("spotify returned non-2xx status")
            w.WriteHeader(http.StatusBadGateway)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status": "error",
                "message": fmt.Sprintf("Spotify API returned %s", resp.Status),
                "body": string(body),
            })
            return
        }

        if err := json.Unmarshal(body, &pageData); err != nil {
            log.Println("decode error:", err)
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  "error",
                "message": "Invalid response from Spotify",
                "detail":  err.Error(),
                "body":    string(body),
            })
            return
        }

        fmt.Printf("%v\n", pageData.Items)

        playlistData = append(playlistData, pageData.Items...)
        url = pageData.Next
    }

	fmt.Printf("Fetched %d tracks from Spotify playlist\n", len(playlistData))

	for i := range playlistData {
		fmt.Printf("%v", playlistData[i])
	}
	fmt.Printf("\n")

	result := map[string]interface{}{
		"status":	"success",
		"tracks":	playlistData,
	}

	json.NewEncoder(w).Encode(result)
}