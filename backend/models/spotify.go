package models

type SpotifyTrackNumberResponse struct {
	Tracks struct {
		Total int `json:"total"`
	} `json:"tracks"`
}

/* 
	This struct will contain 100 items per request, since Spotify's API
	limits the number of items returned in a single request to 100.
	We need to paginate through the results to get all tracks and
	store them in a slice of SpotifyPlaylistItem, which is defined below.
*/
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