package spotify

type Playlist struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}