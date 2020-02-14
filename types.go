package gphotos

// AlbumPosition represents a specified position in an album.
// Source: https://developers.google.com/photos/library/reference/rest/v1/AlbumPosition
type AlbumPosition struct {
	Position                 PositionType `json:"position,omitempty"`
	RelativeMediaItemID      string       `json:"relativeMediaItemId,omitempty"`
	RelativeEnrichmentItemID string       `json:"relativeEnrichmentItemId,omitempty"`
}

// PositionType represents possible positions in an album.
// Source: https://developers.google.com/photos/library/reference/rest/v1/AlbumPosition#positiontype
type PositionType int

// Possible positions in an album.
const (
	// Default value if this enum isn't set.
	PositionTypeUnspecified PositionType = iota

	// At the beginning of the album.
	FirstInAlbum

	// At the end of the album.
	LastInAlbum

	// After a media item.
	AfterMediaItem

	// After an enrichment item.
	AfterEnrichmentItem
)

// Status represents a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs.
// Source: https://developers.google.com/photos/library/reference/rest/v1/Status
type Status struct {
	Code    int               `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}
