package spotify

import (
	"context"
	"time"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type SpotifyClient struct {
	client *spotify.Client
}

func NewSpotifyClient(authToken string) *SpotifyClient {
	authenticator := spotifyauth.New()
	token := &oauth2.Token{AccessToken: authToken}
	client := spotify.New(authenticator.Client(context.Background(), token))

	return &SpotifyClient{client: client}
}

func (sc *SpotifyClient) FetchPlaylists(userID string) ([]spotify.SimplePlaylist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	playlists, err := sc.client.GetPlaylistsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return playlists.Playlists, nil
}

func (sc *SpotifyClient) FetchTracks(playlistID string) ([]spotify.FullTrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := spotify.ID(playlistID)
	tracks, err := sc.client.GetPlaylistTracks(ctx, id)
	if err != nil {
		return nil, err
	}

	var fullTracks []spotify.FullTrack
	for _, track := range tracks.Tracks {
		fullTracks = append(fullTracks, track.Track)
	}

	return fullTracks, nil
}

func (sc *SpotifyClient) GetPlaylist(playlistID string) (*Playlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := spotify.ID(playlistID)
	spotifyPlaylist, err := sc.client.GetPlaylist(ctx, id)
	if err != nil {
		return nil, err
	}

	tracks := make([]Track, 0)
	for _, track := range spotifyPlaylist.Tracks.Tracks {
		if len(track.Track.Artists) > 0 {
			tracks = append(tracks, Track{
				ID:     string(track.Track.ID),
				Name:   track.Track.Name,
				Artist: track.Track.Artists[0].Name,
				Album:  track.Track.Album.Name,
			})
		}
	}

	return &Playlist{
		ID:          string(spotifyPlaylist.ID),
		Name:        spotifyPlaylist.Name,
		Description: spotifyPlaylist.Description,
		Tracks:      tracks,
	}, nil
}
