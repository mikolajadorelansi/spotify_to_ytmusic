package spotify

import (
	"context"
	"net/http"
	"time"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type SpotifyClient struct {
	client *spotify.Client
}

func NewSpotifyClient(authToken string) *SpotifyClient {
	token := &oauth2.Token{AccessToken: authToken}
	httpClient := spotify.NewAuthenticator(http.DefaultClient, token)
	client := spotify.NewClient(httpClient)

	return &SpotifyClient{client: &client}
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

	tracks, err := sc.client.GetPlaylistTracks(ctx, playlistID)
	if err != nil {
		return nil, err
	}

	return tracks.Tracks, nil
}