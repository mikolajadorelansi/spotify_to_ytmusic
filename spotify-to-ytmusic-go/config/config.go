package config

import (
	"os"
)

type Config struct {
	SpotifyAPIKey    string
	YouTubeMusicAPIKey string
}

func LoadConfig() (*Config, error) {
	return &Config{
		SpotifyAPIKey:    os.Getenv("SPOTIFY_API_KEY"),
		YouTubeMusicAPIKey: os.Getenv("YOUTUBE_MUSIC_API_KEY"),
	}, nil
}