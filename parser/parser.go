package parser

import (
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/raulcodes/spotifyWebAPI/types"
	"github.com/raulcodes/test/html"
)

type Playlist struct {
	types.PlaylistObj
}

type PlaylistImpl interface {
	HandlePlaylistRes() (string, error)
}

func NewPlaylistParser(playlist types.PlaylistObj) Playlist {
	return Playlist{playlist}
}

func (p Playlist) HandlePlaylistRes() (html.ParentInfo, error) {
	playlistHtmlStr, err := handlePlaylistHeader(&p.PlaylistObj)
	if err != nil {
		return html.ParentInfo{}, err
	}

	var tracks strings.Builder
	for _, v := range p.Tracks.Items {
		track := v.Track
		if track.Track != nil {
			trackHtmlStr, err := handleTrack(track.Track)
			if err != nil {
				return html.ParentInfo{}, err
			}
			tracks.WriteString(trackHtmlStr)
		}
	}

	result := html.ParentInfo{
		Header: playlistHtmlStr,
		Tracks: tracks.String(),
	}

	return result, nil
}

func (p Playlist) HandleParentTemplate(info html.ParentInfo) error {
	tmpl, err := template.New("html").Parse(html.ParentTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, info)
	if err != nil {
		return err
	}

	return nil
}

func handleTrack(track *types.TrackObj) (string, error) {
	html, err := html.TrackHTML(*track)
	if err != nil {
		return "", err
	}

	err = getAlbumArt(track)
	if err != nil {
		return "", err
	}
	return html, nil
}

func handlePlaylistHeader(playlist *types.PlaylistObj) (string, error) {
	str, err := html.PlaylistHeaderHTML(*playlist)
	err = getPlaylistImg(playlist.Images[0].URL)
	if err != nil {
		return "", err
	}
	return str, nil
}

func getPlaylistImg(url string) error {
	err := downloadFile("images/playlist-cover.png", url)
	if err != nil {
		return err
	}
	return nil
}

func getAlbumArt(track *types.TrackObj) error {
	filepath := html.BuildFilePath(track.ID)

	err := downloadFile(filepath, track.Album.Images[0].URL)
	if err != nil {
		return err
	}

	return nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.MkdirAll("images", 0777)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
