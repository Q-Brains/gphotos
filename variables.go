package gphotos

// Albums is a collection of request methods belonging to `albums`.
// Source: https://developers.google.com/photos/library/reference/rest/v1/albums
var Albums AlbumsRequests = albumsRequests{}

type AlbumsRequests interface {
	baseURL() string
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
