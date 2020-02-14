package gphotos

import (
	"errors"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Uploader is the only instance of UploadMethods.
var Uploader UploadMethods = uploadMethods{}

// UploadMethods is a collection of customized upload methods.
type UploadMethods interface {
	// Upload is a method to upload MediaItems to GooglePhotos.
	// Use UploadWithAlbum or UploadWithAlbumname if you want to add these MediaItems to album at the same time as upload.
	Upload(client *http.Client, filePaths []string) ([]MediaItem, error)

	// UploadWithAlbum is a method to upload MediaItems to GooglePhotos with it added to the Album.
	UploadWithAlbum(client *http.Client, filePaths []string, album Album) ([]MediaItem, error)

	// UploadWithAlbumname is a method to upload MediaItems to GooglePhotos with it added to a specific name Album.
	// If the Album does not exist in GooglePhotos, it will be created.
	UploadWithAlbumname(client *http.Client, filePaths []string, albumname string) (Album, []MediaItem, error)
}

type uploadMethods struct{}

func (uploader uploadMethods) Upload(client *http.Client, filePaths []string) ([]MediaItem, error) {
	req := MediaItemsBatchCreateRequest{}

	err := appendMediaItems(&req, client, filePaths)
	if err != nil {
		return nil, err
	}

	return upload(client, req)
}

func (uploader uploadMethods) UploadWithAlbum(client *http.Client, filePaths []string, album Album) ([]MediaItem, error) {
	req := MediaItemsBatchCreateRequest{
		AlbumID: album.ID,
	}
	if err := appendMediaItems(&req, client, filePaths); err != nil {
		return nil, err
	}

	items, err := upload(client, req)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func appendMediaItems(req *MediaItemsBatchCreateRequest, client *http.Client, filePaths []string) error {
	for _, filePath := range filePaths {
		pathStrs := strings.Split(filePath, "/")
		filename := pathStrs[len(pathStrs)-1]
		token, err := UploadingMedia.UploadMedia(client, filePath, filename)
		if err != nil {
			return err
		}

		req.NewMediaItems = append(
			req.NewMediaItems,
			NewMediaItem{
				Description: filePath,
				SimpleMediaItem: SimpleMediaItem{
					UploadToken: token,
				},
			},
		)
	}
	return nil
}

func upload(client *http.Client, req MediaItemsBatchCreateRequest) ([]MediaItem, error) {
	resp, err := MediaItems.BatchCreate(client, req)
	if err != nil {
		return nil, err
	}

	var items []MediaItem
	for _, result := range resp.NewMediaItemResults {
		if reflect.DeepEqual(result.Status.Message, "OK") {
			if err := os.Remove(result.MediaItem.Description); err != nil {
				return nil, err
			}
			items = append(items, result.MediaItem)
		} else {
			return nil, errors.New("NewMediaItemResult.Status.Message is not \"OK\"")
		}
	}
	return items, nil
}

func (uploader uploadMethods) UploadWithAlbumname(client *http.Client, filePaths []string, albumname string) (Album, []MediaItem, error) {
	album, err := searchAlbum(client, albumname)
	if err != nil {
		return Album{}, nil, err
	}

	items, err := uploader.UploadWithAlbum(client, filePaths, album)
	if err != nil {
		return album, nil, err
	}

	return album, items, nil
}

func searchAlbum(client *http.Client, albumname string) (Album, error) {
	var nextPageToken string
	for true {
		queries := []ListQuery{}
		queries = append(queries, PageSize(1))
		if !reflect.DeepEqual(nextPageToken, "") {
			queries = append(queries, PageToken(nextPageToken))
		}

		resp, err := Albums.List(client, queries...)
		if err != nil {
			return Album{}, err
		}

		for _, album := range resp.Albums {
			if reflect.DeepEqual(album.Title, albumname) {
				return album, nil
			}
		}

		nextPageToken = resp.NextPageToken
		if reflect.DeepEqual(nextPageToken, "") {
			break
		}
	}

	return createAlbum(client, albumname)
}

func createAlbum(client *http.Client, albumname string) (Album, error) {
	req := AlbumsCreateRequest{
		Album: Album{
			Title: albumname,
		},
	}

	resp, err := Albums.Create(client, req)
	if err != nil {
		return Album{}, err
	}

	return Album(resp), nil
}
