package gphotos

import (
	"bytes"
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

// - join

// Join is a method that joins a shared album on behalf of the Google Photos user.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/join
func (sharedAlbums *sharedAlbumsRequests) Join(client *http.Client, request SharedAlbumsJoinRequest) (SharedAlbumsJoinResponse, error) {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return SharedAlbumsJoinResponse{}, err
	}
	req, err := http.NewRequest("POST", sharedAlbums.baseURL()+":join", bytes.NewBuffer(outputJSON))
	resp, err := client.Do(req)
	if err != nil {
		return SharedAlbumsJoinResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return SharedAlbumsJoinResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SharedAlbumsJoinResponse{}, err
	}
	var response SharedAlbumsJoinResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return SharedAlbumsJoinResponse{}, err
	}
	return response, nil
}

// SharedAlbumsJoinRequest is a required body of the SharedAlbums.Join method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/join#request-body
type SharedAlbumsJoinRequest struct {
	ShareToken string `json:"shareToken,omitempty"`
}

// SharedAlbumsJoinResponse is the body returned by the SharedAlbums.Join method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/join#response-body
type SharedAlbumsJoinResponse struct {
	Album Album `json:"album,omitempty"`
}
