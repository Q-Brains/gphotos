package gphotos

import (
	"encoding/json"
)

// Resource: albums

// - Overview

// Album represents an album in Google Photos.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums#resource:-album
type Album struct {
	ID                    string      `json:"id,omitempty"`
	Title                 string      `json:"title,omitempty"`
	ProductURL            string      `json:"productUrl,omitempty"`
	IsWriteable           bool        `json:"isWriteable,omitempty"`
	ShareInfo             ShareInfo   `json:"shareInfo,omitempty"`
	MediaItemsCount       json.Number `json:"mediaItemsCount,omitempty"`
	CoverPhotoBaseURL     string      `json:"coverPhotoBaseUrl,omitempty"`
	CoverPhotoMediaItemID string      `json:"coverPhotoMediaItemId,omitempty"`
}

// ShareInfo represents information about albums that are shared.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums#shareinfo
type ShareInfo struct {
	SharedAlbumOptions SharedAlbumOptions `json:"sharedAlbumOptions,omitempty"`
	ShareableURL       string             `json:"shareableUrl,omitempty"`
	ShareToken         string             `json:"shareToken,omitempty"`
	IsJoined           bool               `json:"isJoined,omitempty"`
	IsOwned            bool               `json:"isOwned,omitempty"`
}

// SharedAlbumOptions represents options that control the sharing of an album.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums#sharedalbumoptions
type SharedAlbumOptions struct {
	IsCollaborative bool `json:"isCollaborative,omitempty"`
	IsCommentable   bool `json:"isCommentable,omitempty"`
}
