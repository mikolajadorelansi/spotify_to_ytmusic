package ytmusic

type YTTrack struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	Duration  int    `json:"duration"` // Duration in seconds
	Thumbnail string `json:"thumbnail"`
}

type YTPlaylist struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tracks      []YTTrack `json:"tracks"`
}
