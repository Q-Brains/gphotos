package gphotos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Albums is the only instance of AlbumsRequests(https://godoc.org/github.com/Q-Brains/gphotos#AlbumsRequests).
var Albums AlbumsRequests = albumsRequests{}

// AlbumsRequests is a collection of requests methods belonging to `albums`.
// The only instance of AlbumsRequests is Albums(https://godoc.org/github.com/Q-Brains/gphotos#Albums).
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums
type AlbumsRequests interface {
	baseURL() string

	// AddEnrichment is a method that adds an enrichment at a specified position in a defined album.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment
	AddEnrichment(client *http.Client, albumID string, request AlbumsAddEnrichmentRequest) (AlbumsAddEnrichmentResponse, error)

	// BatchAddMediaItems is a method that adds one or more media items in a user's Google Photos library to an album.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchAddMediaItems
	BatchAddMediaItems(client *http.Client, albumID string, request AlbumsBatchAddMediaItemsRequest) error

	// BatchRemoveMediaItems is a method that removes one or more media items from a specified album.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchRemoveMediaItems
	BatchRemoveMediaItems(client *http.Client, albumID string, request AlbumsBatchRemoveMediaItemsRequest) error

	// Create is a method that creates an album in a user's Google Photos library.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/create
	Create(client *http.Client, request AlbumsCreateRequest) (AlbumsCreateResponse, error)

	// Get is a method that returns the album based on the specified `albumId`.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/get
	Get(client *http.Client, albumID string) (AlbumsGetResponse, error)

	// List is a method that lists all albums shown to a user in the Albums tab of the Google Photos app.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/list
	List(client *http.Client, queries ...ListQuery) (AlbumsListResponse, error)

	// Share is a method that marks an album as shared and accessible to other users.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/share
	Share(client *http.Client, albumID string, request AlbumsShareRequest) (AlbumsShareResponse, error)

	// Unshare is a method that marks a previously shared album as private.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/unshare
	Unshare(client *http.Client, albumID string) error
}

type albumsRequests struct{}

func (albums albumsRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/albums"
}

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

// - addEnrichment

func (albums albumsRequests) AddEnrichment(client *http.Client, albumID string, request AlbumsAddEnrichmentRequest) (AlbumsAddEnrichmentResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return AlbumsAddEnrichmentResponse{}, err
	}
	req, err := http.NewRequest("POST", albums.baseURL()+"/"+albumID+":addEnrichment", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsAddEnrichmentResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return AlbumsAddEnrichmentResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AlbumsAddEnrichmentResponse{}, err
	}
	var response AlbumsAddEnrichmentResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return AlbumsAddEnrichmentResponse{}, err
	}
	return response, nil
}

// AlbumsAddEnrichmentRequest is a required body of the Albums.AddEnrichment method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#request-body
type AlbumsAddEnrichmentRequest struct {
	NewEnrichmentItem NewEnrichmentItem `json:"newEnrichmentItem,omitempty"`
	AlbumPosition     AlbumPosition     `json:"albumPosition,omitempty"`
}

// AlbumsAddEnrichmentResponse is the body returned by the Albums.AddEnrichment method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#response-body
type AlbumsAddEnrichmentResponse struct {
	EnrichmentItem EnrichmentItem `json:"enrichmentItem,omitempty"`
}

// NewEnrichmentItem represents a new enrichment item to be added to an album.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#newenrichmentitem
type NewEnrichmentItem struct {
	TextEnrichment     TextEnrichment     `json:"textEnrichment,omitempty"`
	LocationEnrichment LocationEnrichment `json:"locationEnrichment,omitempty"`
	MapEnrichment      MapEnrichment      `json:"mapEnrichment,omitempty"`
}

// TextEnrichment represents an enrichment containing text.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#textenrichment
type TextEnrichment struct {
	Text string `json:"text,omitempty"`
}

// LocationEnrichment represents an enrichment containing a single location.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#locationenrichment
type LocationEnrichment struct {
	Location Location `json:"location,omitempty"`
}

// Location represents a physical location.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#location
type Location struct {
	LocationName string `json:"locationName,omitempty"`
	Latlng       LatLng `json:"latlng,omitempty"`
}

// LatLng represents a latitude/longitude pair.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#latlng
type LatLng struct {
	Latitude  int `json:"latitude,omitempty"`
	Longitude int `json:"longitude,omitempty"`
}

// MapEnrichment represents an enrichment containing a map, showing origin and destination locations.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#mapenrichment
type MapEnrichment struct {
	Origin      Location `json:"origin,omitempty"`
	Destination Location `json:"destination,omitempty"`
}

// EnrichmentItem represents an enrichment item.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/addEnrichment#enrichmentitem
type EnrichmentItem struct {
	ID string `json:"id,omitempty"`
}

// - batchAddMediaItems

func (albums albumsRequests) BatchAddMediaItems(client *http.Client, albumID string, request AlbumsBatchAddMediaItemsRequest) error {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", albums.baseURL()+"/"+albumID+":batchAddMediaItems", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return e
	}
	return nil
}

// AlbumsBatchAddMediaItemsRequest is a required body of the Albums.AddMediaItems method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchAddMediaItems#request-body
type AlbumsBatchAddMediaItemsRequest struct {
	MediaItemIDs []string `json:"mediaItemIds,omitempty"`
}

// - batchRemoveMediaItems

func (albums albumsRequests) BatchRemoveMediaItems(client *http.Client, albumID string, request AlbumsBatchRemoveMediaItemsRequest) error {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", albums.baseURL()+"/"+albumID+":batchRemoveMediaItems", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return e
	}
	return nil
}

// AlbumsBatchRemoveMediaItemsRequest is a required body of Albums.BatchRemoveMediaItems method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchRemoveMediaItems#request-body
type AlbumsBatchRemoveMediaItemsRequest struct {
	MediaItemIDs []string `json:"mediaItemIds,omitempty"`
}

// - create

func (albums albumsRequests) Create(client *http.Client, request AlbumsCreateRequest) (AlbumsCreateResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return AlbumsCreateResponse{}, err
	}
	req, err := http.NewRequest("POST", albums.baseURL(), bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsCreateResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return AlbumsCreateResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AlbumsCreateResponse{}, err
	}
	var response AlbumsCreateResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return AlbumsCreateResponse{}, err
	}
	return response, nil
}

// AlbumsCreateRequest is a required body of Albums.Create method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchRemoveMediaItems#request-body
type AlbumsCreateRequest struct {
	Album Album `json:"album,omitempty"`
}

// AlbumsCreateResponse is the body returned by the Albums.Create method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/batchRemoveMediaItems#response-body
type AlbumsCreateResponse Album

// - get

func (albums albumsRequests) Get(client *http.Client, albumID string) (AlbumsGetResponse, error) {
	req, err := http.NewRequest("GET", albums.baseURL()+"/"+albumID, nil)
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsGetResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return AlbumsGetResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AlbumsGetResponse{}, err
	}
	var response AlbumsGetResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return AlbumsGetResponse{}, err
	}
	return response, nil
}

// AlbumsGetResponse is the body returned by the Albums.Get method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/get#response-body
type AlbumsGetResponse Album

// - list

func (albums albumsRequests) List(client *http.Client, queries ...ListQuery) (AlbumsListResponse, error) {
	values := url.Values{}
	for _, query := range queries {
		query(&values)
	}
	req, err := http.NewRequest("GET", albums.baseURL(), nil)
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsListResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return AlbumsListResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AlbumsListResponse{}, err
	}
	var response AlbumsListResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return AlbumsListResponse{}, err
	}
	return response, nil
}

// AlbumsListResponse is the body returned by the Albums.List method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/list#response-body
type AlbumsListResponse struct {
	Albums        []Album `json:"albums,omitempty"`
	NextPageToken string  `json:"nextPageToken,omitempty"`
}

// ListQuery is a structure for using variable length arguments in Albums.List, MediaItems.List and SharedAlbums.List.
type ListQuery func(*url.Values)

// PageSize is a function for passing a page size query to Albums.List, MediaItems.List and SharedAlbums.List.
func PageSize(size int) ListQuery {
	return func(v *url.Values) {
		v.Add("pageSize", strconv.Itoa(size))
	}
}

// PageToken is a function for passing a page token query to Albums.List, MediaItems.List and SharedAlbums.List.
func PageToken(token string) ListQuery {
	return func(v *url.Values) {
		v.Add("pageToken", token)
	}
}

// ExcludeNonAppCreatedData is a function to pass a boolean value to whether to exclude the value created by App to Albums.List, MediaItems.List and SharedAlbums.List.
func ExcludeNonAppCreatedData(flag bool) ListQuery {
	return func(v *url.Values) {
		v.Add("excludeNonAppCreatedData", strconv.FormatBool(flag))
	}
}

// - share

func (albums albumsRequests) Share(client *http.Client, albumID string, request AlbumsShareRequest) (AlbumsShareResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return AlbumsShareResponse{}, err
	}
	req, err := http.NewRequest("POST", albums.baseURL()+"/"+albumID+":share", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsShareResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return AlbumsShareResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AlbumsShareResponse{}, err
	}
	var response AlbumsShareResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return AlbumsShareResponse{}, err
	}
	return response, nil
}

// AlbumsShareRequest is a required body of Albums.Share method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/share#request-body
type AlbumsShareRequest struct {
	SharedAlbumOptions SharedAlbumOptions `json:"sharedAlbumOptions,omitempty"`
}

// AlbumsShareResponse is the body returned by the Albums.Share method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums/share#response-body
type AlbumsShareResponse struct {
	ShareInfo ShareInfo `json:"shareInfo,omitempty"`
}

// - unshare

func (albums albumsRequests) Unshare(client *http.Client, albumID string) error {
	req, err := http.NewRequest("POST", albums.baseURL()+"/"+albumID+":unshare", nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return e
	}
	return nil
}
