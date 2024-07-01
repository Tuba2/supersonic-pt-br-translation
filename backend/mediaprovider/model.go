package mediaprovider

// Bit field flag for the ReleaseTypes property
type ReleaseType = int32

// Bit field of release types
type ReleaseTypes = int32

// Set of possible release types
// Taken from Picard:
//
//	(a) https://picard-docs.musicbrainz.org/en/config/options_releases.html
//	(b) https://musicbrainz.org/doc/Release_Group/Type
const (
	ReleaseTypeAlbum          ReleaseType = 0x0001
	ReleaseTypeAudiobook      ReleaseType = 0x0002
	ReleaseTypeAudioDrama     ReleaseType = 0x0004
	ReleaseTypeBroadcast      ReleaseType = 0x0008
	ReleaseTypeCompilation    ReleaseType = 0x0010
	ReleaseTypeDemo           ReleaseType = 0x0020
	ReleaseTypeDJMix          ReleaseType = 0x0040
	ReleaseTypeEP             ReleaseType = 0x0080
	ReleaseTypeFieldRecording ReleaseType = 0x0100
	ReleaseTypeInterview      ReleaseType = 0x0200
	ReleaseTypeLive           ReleaseType = 0x0400
	ReleaseTypeMixtape        ReleaseType = 0x0800
	ReleaseTypeRemix          ReleaseType = 0x1000
	ReleaseTypeSingle         ReleaseType = 0x2000
	ReleaseTypeSoundtrack     ReleaseType = 0x4000
	ReleaseTypeSpokenWord     ReleaseType = 0x8000
)

type Album struct {
	ID           string
	CoverArtID   string
	Name         string
	Duration     int
	ArtistIDs    []string
	ArtistNames  []string
	Year         int
	ReissueYear  int
	Genres       []string
	TrackCount   int
	Favorite     bool
	ReleaseTypes ReleaseTypes
}

type AlbumWithTracks struct {
	Album
	Tracks []*Track
}

type AlbumInfo struct {
	Notes         string
	LastFmUrl     string
	MusicBrainzID string
}

type Artist struct {
	ID         string
	CoverArtID string
	Name       string
	Favorite   bool
	AlbumCount int
}

type ArtistWithAlbums struct {
	Artist
	Albums []*Album
}

type ArtistInfo struct {
	Biography      string
	LastFMUrl      string
	ImageURL       string
	SimilarArtists []*Artist
}

type Genre struct {
	Name       string
	AlbumCount int
	TrackCount int
}

type Track struct {
	ID          string
	CoverArtID  string
	ParentID    string
	Title       string
	Duration    int
	TrackNumber int
	DiscNumber  int
	Genre       string
	ArtistIDs   []string
	ArtistNames []string
	Album       string
	AlbumID     string
	Year        int
	Rating      int
	Favorite    bool
	Size        int64
	PlayCount   int
	FilePath    string
	BitRate     int
	Comment     string
}

type Playlist struct {
	ID          string
	CoverArtID  string
	Name        string
	Description string
	Public      bool
	Owner       string
	Duration    int
	TrackCount  int
}

type PlaylistWithTracks struct {
	Playlist
	Tracks []*Track
}

type Lyrics struct {
	Title  string
	Artist string
	Synced bool
	Lines  []LyricLine
}

type LyricLine struct {
	Text  string
	Start float64 // seconds
}

type SavedPlayQueue struct {
	Tracks   []*Track
	TrackPos int
	TimePos  int // seconds
}

type RadioStation struct {
	Name        string
	ID          string
	HomePageURL string
	StreamURL   string
}

type MediaItemType int

const (
	MediaItemTypeTrack MediaItemType = iota
	MediaItemTypeRadioStation
)

type MediaItemMetadata struct {
	Type       MediaItemType
	ID         string
	Name       string
	Artists    []string
	ArtistIDs  []string
	Album      string
	AlbumID    string
	CoverArtID string
	Duration   int
}

type MediaItem interface {
	Metadata() MediaItemMetadata
	Copy() MediaItem
}

func (t *Track) Metadata() MediaItemMetadata {
	if t == nil {
		return MediaItemMetadata{}
	}
	return MediaItemMetadata{
		Type:       MediaItemTypeTrack,
		ID:         t.ID,
		Name:       t.Title,
		Artists:    t.ArtistNames,
		ArtistIDs:  t.ArtistIDs,
		Album:      t.Album,
		AlbumID:    t.AlbumID,
		CoverArtID: t.CoverArtID,
		Duration:   t.Duration,
	}
}

func (t *Track) Copy() MediaItem {
	new := *t
	return &new
}

func (r *RadioStation) Metadata() MediaItemMetadata {
	if r == nil {
		return MediaItemMetadata{}
	}
	return MediaItemMetadata{
		Type: MediaItemTypeRadioStation,
		ID:   r.ID,
		Name: r.Name,
	}
}

func (r *RadioStation) Copy() MediaItem {
	return r // no need to copy since RadioStations are immutable
}

type ContentType int

const (
	ContentTypeAlbum ContentType = iota
	ContentTypeArtist
	ContentTypePlaylist
	ContentTypeTrack
	ContentTypeGenre
	ContentTypeRadioStation
)

func (c ContentType) String() string {
	switch c {
	case ContentTypeAlbum:
		return "Album"
	case ContentTypeArtist:
		return "Artist"
	case ContentTypeTrack:
		return "Track"
	case ContentTypePlaylist:
		return "Playlist"
	case ContentTypeGenre:
		return "Genre"
	case ContentTypeRadioStation:
		return "Radio station"
	default:
		return "Unknown"
	}
}

type SearchResult struct {
	Name    string
	ID      string
	CoverID string
	Type    ContentType

	// for Album / Playlist: track count
	//     Artist / Genre: album count
	//     Track: length (seconds)
	// for Radio: none
	Size int

	// Unset for ContentTypes Artist, Playlist, Genre, and RadioStation
	ArtistName string
}
