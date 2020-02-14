package gphotos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

// - batchCreate

// BatchCreate is a method that creates one or more media items in a user's Google Photos library.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate
func (mediaItems *mediaItemsRequests) BatchCreate(client *http.Client, request MediaItemsBatchCreateRequest) (MediaItemsBatchCreateResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return MediaItemsBatchCreateResponse{}, err
	}
	req, err := http.NewRequest("POST", mediaItems.baseURL()+":batchCreate", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return MediaItemsBatchCreateResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return MediaItemsBatchCreateResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MediaItemsBatchCreateResponse{}, err
	}
	var response MediaItemsBatchCreateResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return MediaItemsBatchCreateResponse{}, err
	}
	return response, nil
}

// MediaItemsBatchCreateRequest is a required body of the MediaItems.BatchCreate method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate#request-body
type MediaItemsBatchCreateRequest struct {
	AlbumID       string         `json:"albumId,omitempty"`
	NewMediaItems []NewMediaItem `json:"newMediaItems,omitempty"`
	AlbumPosition AlbumPosition  `json:"albumPosition,omitempty"`
}

// MediaItemsBatchCreateResponse is the body returned by the MediaItems.BatchCreate method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate#response-body
type MediaItemsBatchCreateResponse struct {
	NewMediaItemResults []NewMediaItemResult `json:"newMediaItemResults,omitempty"`
}

// NewMediaItem represents new media item that's created in a user's Google Photos account.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate#newmediaitem
type NewMediaItem struct {
	Description     string          `json:"description,omitempty"`
	SimpleMediaItem SimpleMediaItem `json:"simpleMediaItem,omitempty"`
}

// SimpleMediaItem represents a simple media item to be created in Google Photos via an upload token.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate#simplemediaitem
type SimpleMediaItem struct {
	UploadToken string `json:"uploadToken,omitempty"`
}

// NewMediaItemResult represents result of creating a new media item.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate#newmediaitemresult
type NewMediaItemResult struct {
	UploadToken string    `json:"uploadToken,omitempty"`
	Status      Status    `json:"status,omitempty"`
	MediaItem   MediaItem `json:"mediaItem,omitempty"`
}

// - batchGet

// BatchGet is a method that returns the list of media items for the specified media item identifiers.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchGet
func (mediaItems *mediaItemsRequests) BatchGet(client *http.Client, queries ...MediaItemsBatchGetQuery) (MediaItemsBatchGetResponse, error) {
	values := url.Values{}
	for _, query := range queries {
		query(&values)
	}
	req, err := http.NewRequest("GET", mediaItems.baseURL()+":batchGet", nil)
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return MediaItemsBatchGetResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return MediaItemsBatchGetResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MediaItemsBatchGetResponse{}, err
	}
	var response MediaItemsBatchGetResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return MediaItemsBatchGetResponse{}, err
	}
	return response, nil
}

// MediaItemsBatchGetResponse is the body returned by the MediaItems.BatchGet method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchGet#response-body
type MediaItemsBatchGetResponse struct {
	MediaItemResults []MediaItemResult `json:"mediaItemResults,omitempty"`
}

// MediaItemsBatchGetQuery is a structure for using variable length arguments in MediaItems.BatchGet.
type MediaItemsBatchGetQuery func(*url.Values)

// MediaItemIDs is a function for passing media item indexes query to MediaItems.BatchGet.
func MediaItemIDs(ids string) MediaItemsBatchGetQuery {
	return func(v *url.Values) {
		v.Add("mediaItemIds", ids)
	}
}

// MediaItemResult represents result of retrieving a media item.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchGet#mediaitemresult
type MediaItemResult struct {
	Status    Status    `json:"status,omitempty"`
	MediaItem MediaItem `json:"mediaItem,omitempty"`
}
