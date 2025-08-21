package main

import (
	"fmt"
	"reflect"

	"github.com/zmb3/spotify/v2"
)

func main() {
	var item spotify.PlaylistItem
	fmt.Printf("PlaylistItem fields:\n")
	t := reflect.TypeOf(item)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: %v\n", field.Name, field.Type)
	}

	fmt.Printf("\nPlaylistTrack fields:\n")
	var track spotify.PlaylistTrack
	t = reflect.TypeOf(track)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: %v\n", field.Name, field.Type)
	}
}
