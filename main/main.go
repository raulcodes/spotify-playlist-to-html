package main

import (
	"os"

	"github.com/raulcodes/test/parser"

	"github.com/raulcodes/spotifyWebAPI/authorization"
	"github.com/raulcodes/spotifyWebAPI/playlist"
)

func main() {
	arg := os.Args[1]

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	client := authorization.NewClient(clientID, clientSecret)
	tokenRes, err := client.AccessToken()
	if err != nil {
		panic(err)
	}

	playlistClient := playlist.NewClient(tokenRes.AccessToken)
	playlist, err := playlistClient.GetPlaylist(arg)
	if err != nil {
		panic(err)
	}

	parser := parser.NewPlaylistParser(playlist)
	parent, err := parser.HandlePlaylistRes()
	if err != nil {
		panic(err)
	}

	err = parser.HandleParentTemplate(parent)
	if err != nil {
		panic(err)
	}
}
