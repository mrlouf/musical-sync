package models

type DeezerTrackNumberResponse struct {
	Total int `json:"nb_tracks"`
}

type DeezerTracklistResponse struct {
	Title string `json:"title"`
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

type DeezerAlbumResponse struct {
	Title string `json:"title"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
	Tracks struct {
		Data []struct {
			Title string `json:"title"`
		} `json:"data"`
	} `json:"tracks"`
}