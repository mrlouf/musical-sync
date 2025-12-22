package models

type DeezerTrackNumberResponse struct {
	Total int `json:"nb_tracks"`
}

type SpotifyTrackNumberResponse struct {
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
}

type DeezerTracklistResponse struct {
	Data []struct {
		Title string `json:"title"`
		Artist struct {
			Name string `json:"name"`
		} `json:"artist"`
		Album struct {
			Title string `json:"title"`
		} `json:"album"`
	} `json:"data"`
}

type DeezerTrackResponse struct {
	Title string `json:"title"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
	Album struct {
		Title string `json:"title"`
	} `json:"album"`
}