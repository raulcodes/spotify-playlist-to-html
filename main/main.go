package main

import (
	"os"

	"github.com/raulcodes/test/parser"

	spotify "github.com/raulcodes/spotify-web-api"
)

func main() {
	arg := os.Args[1]

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	client := spotify.NewClient(clientID, clientSecret)
	tokenRes, err := client.AccessToken()
	if err != nil {
		panic(err)
	}

	client.SetToken(tokenRes.AccessToken)
	playlist, err := client.GetPlaylist(arg)
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