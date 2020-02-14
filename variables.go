package gphotos

import "net/http"

// Albums is the only instance of AlbumsRequests.
var Albums AlbumsRequests = albumsRequests{}

// AlbumsRequests is a collection of requests methods belonging to `albums`.
// The only instance of AlbumsRequests is Albums.
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

// MediaItems is a collection of request methods belonging to `mediaItems`.
// Source: https://developers.google.com/photos/library/reference/rest/v1/mediaItems
var MediaItems MediaItemsRequests = mediaItemsRequests{}

type MediaItemsRequests interface {
	baseURL() string
}

type mediaItemsRequests struct{}

func (mediaItems mediaItemsRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/mediaItems"
}

// SharedAlbums is a collection of request methods belonging to `sharedAlbums`.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums
var SharedAlbums SharedAlbumsRequests = sharedAlbumsRequests{}

type SharedAlbumsRequests interface {
	baseURL() string
}

type sharedAlbumsRequests struct{}

func (sharedAlbums sharedAlbumsRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/sharedAlbums"
}

// Uploader is a collection of customized upload methods.
var Uploader UploadMethods = uploadMethods{}

type UploadMethods interface{}

type uploadMethods struct{}