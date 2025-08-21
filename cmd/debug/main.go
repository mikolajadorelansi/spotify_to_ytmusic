package main

import (
	"fmt"

	"github.com/zmb3/spotify/v2"
)

func main() {
	// Print the type information for PlaylistItem
	pi := spotify.PlaylistItem{}
	fmt.Printf("PlaylistItem type info:\n")
	fmt.Printf("Track field type: %T\n", pi.Track)

	// Print the type information for SimpleTrack
	st := spotify.SimpleTrack{}
	fmt.Printf("\nSimpleTrack type info:\n")
	fmt.Printf("ID type: %T\n", st.ID)
	fmt.Printf("Artists type: %T\n", st.Artists)

	// Print the type information for FullTrack
	ft := spotify.FullTrack{}
	fmt.Printf("\nFullTrack type info:\n")
	fmt.Printf("ID type: %T\n", ft.ID)
	fmt.Printf("Artists type: %T\n", ft.Artists)
	fmt.Printf("Album type: %T\n", ft.Album)
}
