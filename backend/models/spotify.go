package models

type SpotifyTrackNumberResponse struct {
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
}

