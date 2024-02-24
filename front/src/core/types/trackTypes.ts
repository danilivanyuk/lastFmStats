export interface TrackType {
    artist: IArtist
    streamable: string
    image: [IImage]
	mbid: string
	album: IAlbum
	name: string
	"@attr": {
		nowplaying: string
	}
	url: string
	date: {
		uts: string,
		"#text": string
	}
}

export interface IRecentTracks {
	recenttracks: {
		track: [TrackType]
	}
}

interface IImage {
            size: string
            "#text": string
}


interface IArtist {
	mbid: string
        "#text": string
}

interface IAlbum {
	mbid: string
        "#text": string
}

