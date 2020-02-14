package gphotos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MediaItems is the only instance of MediaItemsRequests.
var MediaItems MediaItemsRequests = mediaItemsRequests{}

// MediaItemsRequests is a collection of request methods belonging to `mediaItems`.
// The only instance of MediaItemsRequests is MediaItems.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems
type MediaItemsRequests interface {
	baseURL() string

	// BatchCreate is a method that creates one or more media items in a user's Google Photos library.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate
	BatchCreate(client *http.Client, request MediaItemsBatchCreateRequest) (MediaItemsBatchCreateResponse, error)

	// BatchGet is a method that returns the list of media items for the specified media item identifiers.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchGet
	BatchGet(client *http.Client, queries ...MediaItemsBatchGetQuery) (MediaItemsBatchGetResponse, error)

	// Get is a method that returns the media item for the specified media item identifier.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/get
	Get(client *http.Client, mediaItemID string) (MediaItemsGetResponse, error)

	// List is a method that list all media items from a user's Google Photos library.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/list
	List(client *http.Client, queries ...ListQuery) (MediaItemsListResponse, error)

	// Search is a method that searches for media items in a user's Google Photos library.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search
	Search(client *http.Client, request MediaItemsSearchRequest) (MediaItemsSearchResponse, error)
}

type mediaItemsRequests struct{}

func (mediaItems mediaItemsRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/mediaItems"
}

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

func (mediaItems mediaItemsRequests) BatchCreate(client *http.Client, request MediaItemsBatchCreateRequest) (MediaItemsBatchCreateResponse, error) {
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

func (mediaItems mediaItemsRequests) BatchGet(client *http.Client, queries ...MediaItemsBatchGetQuery) (MediaItemsBatchGetResponse, error) {
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

// - get

func (mediaItems mediaItemsRequests) Get(client *http.Client, mediaItemID string) (MediaItemsGetResponse, error) {
	req, err := http.NewRequest("GET", mediaItems.baseURL()+"/"+mediaItemID, nil)
	resp, err := client.Do(req)
	if err != nil {
		return MediaItemsGetResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return MediaItemsGetResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MediaItemsGetResponse{}, err
	}
	var response MediaItemsGetResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return MediaItemsGetResponse{}, err
	}
	return response, nil
}

// MediaItemsGetResponse is the body returned by the MediaItems.Get method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/get#response-body
type MediaItemsGetResponse MediaItem

// - list

func (mediaItems mediaItemsRequests) List(client *http.Client, queries ...ListQuery) (MediaItemsListResponse, error) {
	values := url.Values{}
	for _, query := range queries {
		query(&values)
	}
	req, err := http.NewRequest("GET", mediaItems.baseURL(), nil)
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return MediaItemsListResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return MediaItemsListResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MediaItemsListResponse{}, err
	}
	var response MediaItemsListResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return MediaItemsListResponse{}, err
	}
	return response, nil
}

// MediaItemsListResponse is the body returned by the MediaItems.BatchGet method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/list#response-body
type MediaItemsListResponse struct {
	MediaItems    []MediaItem `json:"mediaItems,omitempty"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
}

// - search

func (mediaItems mediaItemsRequests) Search(client *http.Client, request MediaItemsSearchRequest) (MediaItemsSearchResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return MediaItemsSearchResponse{}, err
	}
	req, err := http.NewRequest("POST", mediaItems.baseURL()+":search", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return MediaItemsSearchResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return MediaItemsSearchResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MediaItemsSearchResponse{}, err
	}
	var response MediaItemsSearchResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return MediaItemsSearchResponse{}, err
	}
	return response, nil
}

// MediaItemsSearchRequest is a required body of the MediaItems.Search method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#request-body
type MediaItemsSearchRequest struct {
	AlbumID   string  `json:"albumId,omitempty"`
	PageSize  int     `json:"pageSize,omitempty"`
	PageToken string  `json:"pageToken,omitempty"`
	Filters   Filters `json:"filters,omitempty"`
}

// MediaItemsSearchResponse is the body returned by the MediaItems.BatchGet method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#response-body
type MediaItemsSearchResponse struct {
	MediaItems    []MediaItem `json:"mediaItems,omitempty"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
}

// Filters that can be applied to a media item search.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#filters
type Filters struct {
	DateFilter               DateFilter      `json:"dateFilter,omitempty"`
	ContentFilter            ContentFilter   `json:"contentFilter,omitempty"`
	MediaTypeFilter          MediaTypeFilter `json:"mediaTypeFilter,omitempty"`
	FeatureFilter            FeatureFilter   `json:"featureFilter,omitempty"`
	IncludeArchivedMedia     bool            `json:"includeArchivedMedia,omitempty"`
	ExcludeNonAppCreatedData bool            `json:"excludeNonAppCreatedData,omitempty"`
}

// DateFilter represents the allowed dates or date ranges for the media returned.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#datefilter
type DateFilter struct {
	Dates  []Date      `json:"dates,omitempty"`
	Ranges []DateRange `json:"ranges,omitempty"`
}

// Date represents a whole calender date.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#date
type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

// DateRange represents a range of dates.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#daterange
type DateRange struct {
	StartDate Date `json:"startDate,omitempty"`
	EndDate   Date `json:"endDate,omitempty"`
}

// ContentFilter represents the media item returned based on the content type.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#contentfilter
type ContentFilter struct {
	IncludedContentCategories ContentCategory `json:"includedContentCategories,omitempty"`
	ExcludedContentCategories ContentCategory `json:"excludedContentCategories,omitempty"`
}

// ContentCategory represents a set of pre-defined content categories that you can filter on.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#contentcategory
type ContentCategory int

// This is a set of pre-defined content categories that you can filter on.
const (
	// Default content category. This category is ignored when any other category is used in the filter.
	ContentCategoryNone ContentCategory = iota

	// Media items containing landscapes.
	Landscapes

	// Media items containing receipts.
	Receipts

	// Media items containing cityscapes.
	Cityscapes

	// Media items containing landmarks.
	Landmarks

	// Media items that are selfies.
	Selfies

	// Media items containing people.
	People

	// Media items containing pets.
	Pets

	// Media items from weddings.
	Weddings

	// Media items from birthdays.
	Birthdays

	// Media items containing documents.
	Documents

	// Media items taken during travel.
	Travel

	// Media items containing animals.
	Animals

	// Media items containing food.
	Food

	// Media items from sporting events.
	Sport

	// Media items taken at night.
	Night

	// Media items from performances.
	Performances

	// Media items containing whiteboards.
	Whiteboards

	// Media items that are screenshots.
	Screenshots

	// Media items that are considered to be utility.
	// These include, but aren't limited to documents, screenshots, whiteboards etc.
	Utility

	// Media items containing art.
	Arts

	// Media items containing crafts.
	Crafts

	// Media items related to fashion.
	Fashion

	// Media items containing houses.
	Houses

	// Media items containing gardens.
	Gardens

	// Media items containing flowers.
	Flowers

	// Media items taken of holidays.
	Holidays
)

// MediaTypeFilter represents the type of media items to be returned, for example, videos or photos.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#mediatypefilter
type MediaTypeFilter struct {
	MediaTypes MediaType `json:"mediaTypes,omitempty"`
}

// MediaType represents the set of media types that can be searched for.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#mediatype
type MediaType int

// The set of media types that can be searched for.
const (
	// Treated as if no filters are applied. All media types are included.
	AllMedia MediaType = iota

	// All media items that are considered videos.
	// This also includes movies the user has created using the Google Photos app.
	VideoType

	// All media items that are considered photos.
	// This includes .bmp, .gif, .ico, .jpg (and other spellings), .tiff, .webp and special photo types such as iOS live photos, Android motion photos, panoramas, photospheres.
	PhotoType
)

// FeatureFilter represents the features that the media items should have.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#featurefilter
type FeatureFilter struct {
	IncludedFeatures Feature `json:"includedFeatures,omitempty"`
}

// Feature represents the set of features that you can filter on.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/search#feature
type Feature int

// The set of features that you can filter on.
const (
	// Treated as if no filters are applied. All features are included.
	FeatureNone Feature = iota

	// Media items that the user has marked as favorites in the Google Photos app.
	Favorites
)
