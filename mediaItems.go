package gphotos

// Resource: mediaItems

// - Overview

// MediaItem represents a media item (such as a photo or video) in Google Photos.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#resource:-mediaitem
type MediaItem struct {
	ID              string          `json:"id,omitempty"`
	Description     string          `json:"description,omitempty"`
	ProductURL      string          `json:"productUrl,omitempty"`
	BaseURL         string          `json:"baseUrl,omitempty"`
	MimeType        string          `json:"mimeType,omitempty"`
	MediaMetadata   MediaMetadata   `json:"mediaMetadata,omitempty"`
	ContributorInfo ContributorInfo `json:"contributorInfo,omitempty"`
	Filename        string          `json:"filename,omitempty"`
}

// MediaMetadata represents metadata for a media item.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#mediametadata
type MediaMetadata struct {
	CreationTime string `json:"creationTime,omitempty"`
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	Photo        Photo  `json:"photo,omitempty"`
	Video        Video  `json:"video,omitempty"`
}

// Photo represents metadata that is specific to a photo, such as, ISO, focal length and exposure time.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#photo
type Photo struct {
	CameraMake      string `json:"cameraMake,omitempty"`
	CameraModel     string `json:"cameraModel,omitempty"`
	FocalLength     int    `json:"focalLength,omitempty"`
	ApertureFNumber int    `json:"apertureFNumber,omitempty"`
	ISOEquivalent   int    `json:"isoEquivalent,omitempty"`
	ExposureTime    string `json:"exposureTime,omitempty"`
}

// Video represents metadata that is specific to a video, for example, fps and processing status.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#video
type Video struct {
	CameraMake  string                `json:"cameraMake,omitempty"`
	CameraModel string                `json:"cameraModel,omitempty"`
	FPS         int                   `json:"fps,omitempty"`
	Status      VideoProcessingStatus `json:"status,omitempty"`
}

// VideoProcessingStatus represents processing status of a video being uploaded to Google Photos.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#videoprocessingstatus
type VideoProcessingStatus int

const (
	unspecified VideoProcessingStatus = iota
	processing
	ready
	failed
)

// ContributorInfo represents information about the user who added the media item.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems#contributorinfo
type ContributorInfo struct {
	ProfilePictureBaseURL string `json:"profilePictureBaseUrl,omitempty"`
	DisplayName           string `json:"displayName,omitempty"`
}
