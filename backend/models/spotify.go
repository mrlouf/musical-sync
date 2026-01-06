package models

type SpotifyTrackNumberResponse struct {
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
}

type SpotifyPlaylistItem struct {
    Track struct {
        Name    string `json:"name"`
        Artists []struct {
            Name string `json:"name"`
        } `json:"artists"`
        Album struct {
            Name string `json:"name"`
        } `json:"album"`
    } `json:"track"`
}

type SpotifyPlaylistResponse struct {
    Items []SpotifyPlaylistItem `json:"items"`
}