package html

var ParentTemplate = `
	<link rel="stylesheet" media="all" href="style.css">
	<div class="spotify-playlist-container">
		{{ .Header }}
		<div class="spotify-playlist-list">
			{{ .Tracks }}
		</div>
	</div>
`

type ParentInfo struct {
	Header string
	Tracks string
}

var TrackTemplate = `
	<div class="spotify-playlist-track">
		<img class="spotify-playlist-track-img" src="{{ .Img }}" height=100 width=100 />
		<div class="spotify-playlist-track-content">
			<h3 class="spotify-playlist-track-name">{{ .TrackName }}</h3>
			<p class="spotify-playlist-track-artist">{{ .ArtistsNames }}</p>
		</div>
	</div>
`

type TrackInfo struct {
	Img          string
	TrackName    string
	ArtistsNames string
}

var PlaylistHeaderTemplate = `
	<header class="spotify-playlist-header">
		<img class="spotify-playlist-img" src="{{ .Img }}" height=200 width=200 />
		<div class="spotify-playlist-content">
			<h2 class="spotify-playlist-name">{{ .Name }}</h2>
			<p class="spotify-playlist-description>{{ .Description }}</p>
		</div>
	</header>
`

type PlaylistInfo struct {
	Img         string
	Name        string
	Description string
}
