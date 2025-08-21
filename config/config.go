package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration settings
type Config struct {
	SpotifyAPIKey      string `yaml:"spotify_api_key"`
	YouTubeMusicAPIKey string `yaml:"youtube_music_api_key"`
}

// LoadConfig attempts to load configuration from config.yaml
// Falls back to environment variables if file is not found or values are missing
func LoadConfig() (*Config, error) {
	config := &Config{}

	// Try to load from YAML file first
	configPath := filepath.Join("config.yaml")
	if yamlFile, err := os.ReadFile(configPath); err == nil {
		if err := yaml.Unmarshal(yamlFile, config); err != nil {
			return nil, err
		}
	}

	// Fall back to environment variables for any empty values
	if config.SpotifyAPIKey == "" {
		config.SpotifyAPIKey = os.Getenv("SPOTIFY_API_KEY")
	}
	if config.YouTubeMusicAPIKey == "" {
		config.YouTubeMusicAPIKey = os.Getenv("YOUTUBE_MUSIC_API_KEY")
	}

	return config, nil
}
