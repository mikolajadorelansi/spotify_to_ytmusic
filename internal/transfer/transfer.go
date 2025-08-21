package transfer

import (
	"fmt"
	"log"

	"github.com/yourusername/spotify-to-ytmusic-go/internal/spotify"
	"github.com/yourusername/spotify-to-ytmusic-go/internal/ytmusic"
)

type TransferService struct {
	SpotifyClient  *spotify.SpotifyClient
	YouTubeClient  *ytmusic.YouTubeMusicClient
}

func NewTransferService(spotifyClient *spotify.SpotifyClient, youTubeClient *ytmusic.YouTubeMusicClient) *TransferService {
	return &TransferService{
		SpotifyClient: spotifyClient,
		YouTubeClient: youTubeClient,
	}
}

func (t *TransferService) TransferPlaylist(playlistID string) error {
	playlist, err := t.SpotifyClient.GetPlaylist(playlistID)
	if err != nil {
		return fmt.Errorf("failed to fetch playlist from Spotify: %w", err)
	}

	ytPlaylist := t.createYouTubePlaylist(playlist)
	if err := t.YouTubeClient.CreatePlaylist(ytPlaylist); err != nil {
		return fmt.Errorf("failed to create playlist on YouTube Music: %w", err)
	}

	for _, track := range playlist.Tracks {
		ytTrack, err := t.YouTubeClient.SearchTrack(track.Name, track.Artists)
		if err != nil {
			log.Printf("could not find track %s by %s on YouTube Music: %v", track.Name, track.Artists, err)
			continue
		}
		if err := t.YouTubeClient.AddTrackToPlaylist(ytPlaylist.ID, ytTrack.ID); err != nil {
			log.Printf("could not add track %s to YouTube Music playlist: %v", track.Name, err)
		}
	}

	return nil
}

func (t *TransferService) createYouTubePlaylist(playlist *spotify.Playlist) *ytmusic.YTPlaylist {
	return &ytmusic.YTPlaylist{
		Title:       playlist.Name,
		Description: playlist.Description,
	}
}