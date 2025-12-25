package models

type SpotifyTrackNumberResponse struct {
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
}

type SpotifyPlaylistResponse struct {
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