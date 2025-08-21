# Spotify to YouTube Music Transfer

This project is a Go application that allows users to transfer their playlists from Spotify to YouTube Music. It handles authentication with both services and facilitates the migration of playlists and tracks.

## Project Structure

```
spotify-to-ytmusic-go
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── spotify
│   │   ├── client.go    # Spotify API client for authentication and data fetching
│   │   └── models.go    # Data structures for Spotify playlists and tracks
│   ├── ytmusic
│   │   ├── client.go    # YouTube Music API client for authentication and data handling
│   │   └── models.go    # Data structures for YouTube Music playlists and tracks
│   ├── transfer
│       └── transfer.go  # Service for transferring playlists between Spotify and YouTube Music
├── config
│   └── config.go        # Configuration settings and environment variable loading
├── go.mod               # Module definition and dependencies
├── go.sum               # Checksums for module dependencies
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/mikolajadorelansi/spotify_to_ytmusic.git
   cd spotify_to_ytmusic
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Configure API credentials:**
   There are two ways to configure your Spotify and YouTube Music API credentials:

   a. Using a configuration file (recommended):
      - Copy `config.yaml.example` to `config.yaml`
      - Fill in your API keys in the `config.yaml` file:
        ```yaml
        spotify_api_key: "your-spotify-api-key"
        youtube_music_api_key: "your-youtube-music-api-key"
        ```

   b. Using environment variables:
      Set the following environment variables:
      ```
      SPOTIFY_API_KEY=your-spotify-api-key
      YOUTUBE_MUSIC_API_KEY=your-youtube-music-api-key
      ```
      
   The application will first look for a `config.yaml` file and fall back to environment variables if the file is not found or values are missing.

4. **Run the application:**
   Use the following command to start the application:
   ```
   go run cmd/main.go
   ```

## Usage Guidelines

- Follow the prompts to authenticate with your Spotify and YouTube Music accounts.
- Select the playlist you wish to transfer from Spotify.
- The application will search for matching tracks on YouTube Music and create a new playlist with those tracks.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
