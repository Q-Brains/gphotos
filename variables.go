package gphotos

import (
	"net/http"

	"golang.org/x/oauth2"
)

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

// SharedAlbums is the only instance of SharedAlbumsRequests.
var SharedAlbums SharedAlbumsRequests = sharedAlbumsRequests{}

// SharedAlbumsRequests is a collection of request methods belonging to `sharedAlbums`.
// The only instance of SharedAlbumsRequests is SharedAlbums.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums
type SharedAlbumsRequests interface {
	baseURL() string

	// Get is a method that returns the album based on the specified `shareToken`.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/get
	Get(client *http.Client, shareToken string) (SharedAlbumsGetResponse, error)

	// Join is a method that joins a shared album on behalf of the Google Photos user.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/join
	Join(client *http.Client, request SharedAlbumsJoinRequest) (SharedAlbumsJoinResponse, error)

	// Leave is a method that leaves a previously-joined shared album on behalf of the Google Photos user.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/leave
	Leave(client *http.Client, request SharedAlbumsLeaveRequest) error

	// List is a method that lists all shared albums available in the Sharing tab of the user's Google Photos app.
	// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/list
	List(client *http.Client, queries ...ListQuery) (SharedAlbumsListResponse, error)
}

type sharedAlbumsRequests struct{}

func (sharedAlbums sharedAlbumsRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/sharedAlbums"
}

// Auth is the only instance of Authorizations.
var Auth Authorizations = authorizations{}

// Authorizations is a collection of authentication methods.
// However, PhotosLibraryAPI can only be authenticated with OAuth2 authentication.
type Authorizations interface {
	// OAuth2InteractiveFlow is a method that performs OAuth2 authentication of GooglePhotos interactivity.
	// Because it uses standard input, it cannot be used in automation systems.
	OAuth2InteractiveFlow(clientID string, clientSecret string, scopes AuthorizationScopes, state string, options ...oauth2.AuthCodeOption) (*http.Client, error)

	// OAuth2Config is a method that creates the config of "golang.org/x/oauth2".
	// Using this method eliminates the need to import "golang.org/x/oauth2".
	OAuth2Config(clientID string, clientSecret string, scopes AuthorizationScopes) oauth2.Config

	// OAuth2CreateURL is a method that creates OAuth2 authorization URL.
	OAuth2CreateURL(conf oauth2.Config, state string, options ...oauth2.AuthCodeOption) string

	// OAuth2CreateClient is a method that creates OAuth2 client from the config of "golang.org/x/oauth2" and authorization code.
	OAuth2CreateClient(conf oauth2.Config, authCode string) (*http.Client, error)
}

type authorizations struct{}

// UploadingMedia is the only instance of UploadingMediaRequests.
var UploadingMedia UploadingMediaRequests = uploadingMediaRequests{}

// UploadingMediaRequests is a collection of request methods belonging to `UploadingMedia`.
// Source: https://developers.google.com/photos/library/guides/overview
type UploadingMediaRequests interface {
	baseURL() string

	// UploadMedia is a method that uploads media items to a user’s library or album.
	// Source: https://developers.google.com/photos/library/guides/upload-media
	UploadMedia(client *http.Client, filePath string, filename string) (uploadToken string, err error)

	// ResumableUploads is a method.
	// Source: https://developers.google.com/photos/library/guides/resumable-uploads
	ResumableUploads(client *http.Client, filePath string, filename string) (uploadToken string, err error)
}

type uploadingMediaRequests struct{}

func (upload uploadingMediaRequests) baseURL() string {
	return "https://photoslibrary.googleapis.com/v1/uploads"
}

// Uploader is a collection of customized upload methods.
var Uploader UploadMethods = uploadMethods{}

type UploadMethods interface{}

type uploadMethods struct{}
