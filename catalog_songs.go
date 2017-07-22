package applemusic

import (
	"context"
	"fmt"
	"strings"
)

type SongAttributes struct {
	Artwork          Artwork         `json:"artwork"`
	ArtistName       string          `json:"artistName"`
	URL              string          `json:"url"`
	DiscNumber       int             `json:"discNumber"`
	GenreNames       []string        `json:"genreNames"`
	DurationInMillis int64           `json:"durationInMillis,omitempty"`
	ReleaseDate      string          `json:"releaseDate"`
	Name             string          `json:"name"`
	PlayParams       *PlayParameters `json:"playParams,omitempty"`
	TrackNumber      int             `json:"trackNumber,omitempty"`
	ComposerName     string          `json:"composerName,omitempty"`
	ContentRating    string          `json:"contentRating,omitempty"`
	EditorialNotes   *EditorialNotes `json:"editorialNotes,omitempty"`
	MovementCount    int             `json:"movementCount,omitempty"`
	MovementName     string          `json:"movementName,omitempty"`
	MovementNumber   int             `json:"movementNumber,omitempty"`
	WorkName         string          `json:"workName,omitempty"`
}

type SongRelationships struct {
	Albums  Albums  `json:"albums"`           // Default inclusion: Identifiers only
	Artists Artists `json:"artists"`          // Default inclusion: Identifiers only
	Genres  Genres  `json:"genres,omitempty"` // Default inclusion: None
}

// Song represents a song.
type Song struct {
	Id            string            `json:"id"`
	Type          string            `json:"type"`
	Href          string            `json:"href"`
	Attributes    SongAttributes    `json:"attributes"`
	Relationships SongRelationships `json:"relationships"`
}

type Songs struct {
	Data []Song `json:"data"`
	Href string `json:"href,omitempty"`
	Next string `json:"next,omitempty"`
}

func (s *CatalogService) getSongs(ctx context.Context, u string) (*Songs, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	songs := &Songs{}
	resp, err := s.client.Do(ctx, req, songs)
	if err != nil {
		return nil, resp, err
	}

	return songs, resp, nil
}

// GetSong fetches a song using its identifier.
func (s *CatalogService) GetSong(ctx context.Context, storefront, id string, opt *Options) (*Songs, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/songs/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getSongs(ctx, u)
}

// GetSongsByIds fetches one or more songs using their identifiers.
func (s *CatalogService) GetSongsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Songs, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/songs?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getSongs(ctx, u)
}