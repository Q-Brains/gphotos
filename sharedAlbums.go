package gphotos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Resource: sharedAlbums

// - Overview

// - get

// Get is a method that returns the album based on the specified `shareToken`.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/get
func (sharedAlbums *sharedAlbumsRequests) Get(client *http.Client, shareToken string) (SharedAlbumsGetResponse, error) {
	req, err := http.NewRequest("GET", sharedAlbums.baseURL()+"/"+shareToken, nil)
	if err != nil {
		return SharedAlbumsGetResponse{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return SharedAlbumsGetResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return SharedAlbumsGetResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SharedAlbumsGetResponse{}, err
	}
	var response SharedAlbumsGetResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return SharedAlbumsGetResponse{}, err
	}
	return response, nil
}

// SharedAlbumsGetResponse is the body returned by the SharedAlbums.Get method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/get#response-body
type SharedAlbumsGetResponse Album
