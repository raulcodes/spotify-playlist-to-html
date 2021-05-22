package html

var ParentTemplate = `
	<link rel="stylesheet" media="all" href="style.css">
	<div class="spotify-playlist-container">
		{{ .Header }}
	<ul class="spotify-playlist-list">
		{{ .Tracks }}
	</ul>
	</div>`

type ParentInfo struct {
	Header string
	Tracks string
}

var TrackTemplate = `
	<li class="spotify-playlist-track">
		<a href="{{ .Href }}" target="_blank" rel="noreferrer noopener">
			<img class="spotify-playlist-track-img" src="{{ .Img }}" height=100 width=100 />
			<div class="spotify-playlist-track-content">
				<h3 class="spotify-playlist-track-name">{{ .TrackName }}</h3>
				<p class="spotify-playlist-track-artist">{{ .ArtistsNames }}</p>
			</div>
		</a>
	</li>`

type TrackInfo struct {
	Img          string
	TrackName    string
	ArtistsNames string
	Href         string
}

var PlaylistHeaderTemplate = `
	<header class="spotify-playlist-header">
		<a href="{{ .Href }}" target="_blank" rel="noreferrer noopener">
			<img class="spotify-playlist-img" src="{{ .Img }}" height=200 width=200 />
			<div class="spotify-playlist-content">
				<h2 class="spotify-playlist-name">{{ .Name }}</h2>
				<p class="spotify-playlist-description">{{ .Description }}</p>
			</div>
		</a>
	</header>`

type PlaylistInfo struct {
	Img         string
	Name        string
	Description string
	Href        string
}
