package html

import (
	"fmt"
	"strings"
	"text/template"

	spotify "github.com/raulcodes/spotify-web-api"
)

func TrackHTML(track spotify.TrackObj) (string, error) {
	var artists []string
	for _, s := range track.Artists {
		artists = append(artists, s.Name)
	}
	artistStr := strings.Join(artists, ", ")
	img := BuildFilePath(track.ID)

	info := TrackInfo{
		Img:          img,
		TrackName:    track.Name,
		ArtistsNames: artistStr,
	}
	tmpl, err := template.New("html").Parse(TrackTemplate)
	if err != nil {
		return "", err
	}

	var output strings.Builder
	err = tmpl.Execute(&output, info)
	if err != nil {
		return "", nil
	}

	return output.String(), nil
}

func PlaylistHeaderHTML(playlist spotify.PlaylistObj) (string, error) {
	info := PlaylistInfo{
		Img:         "images/playlist-cover.png",
		Name:        playlist.Name,
		Description: playlist.Description,
	}
	tmpl, err := template.New("html").Parse(PlaylistHeaderTemplate)
	if err != nil {
		return "", err
	}

	var output strings.Builder
	err = tmpl.Execute(&output, info)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func BuildFilePath(id string) string {
	filename := strings.ReplaceAll(id, " ", "-")
	return fmt.Sprintf("images/%s.png", filename)
}
