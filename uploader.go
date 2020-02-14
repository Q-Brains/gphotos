package gphotos

import (
	"errors"
	"net/http"
	"os"
	"reflect"
	"strings"
)

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
