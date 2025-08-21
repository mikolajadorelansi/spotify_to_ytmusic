module spotify-to-ytmusic-go

go 1.18

require (
    github.com/go-resty/resty/v2 v2.6.0
    github.com/spf13/viper v1.10.1
)

replace github.com/yourusername/spotify-to-ytmusic-go/internal/spotify => ./internal/spotify
replace github.com/yourusername/spotify-to-ytmusic-go/internal/ytmusic => ./internal/ytmusic
