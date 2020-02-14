package gphotos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SharedAlbums is the only instance of SharedAlbumsRequests(https://godoc.org/github.com/Q-Brains/gphotos#SharedAlbumsRequests).
var SharedAlbums SharedAlbumsRequests = sharedAlbumsRequests{}

// SharedAlbumsRequests is a collection of request methods belonging to `sharedAlbums`.
// The only instance of SharedAlbumsRequests is SharedAlbums(https://godoc.org/github.com/Q-Brains/gphotos#SharedAlbums).
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

// Resource: sharedAlbums

// - Overview

// - get

func (sharedAlbums sharedAlbumsRequests) Get(client *http.Client, shareToken string) (SharedAlbumsGetResponse, error) {
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

func (sharedAlbums sharedAlbumsRequests) Join(client *http.Client, request SharedAlbumsJoinRequest) (SharedAlbumsJoinResponse, error) {
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

// - leave

func (sharedAlbums sharedAlbumsRequests) Leave(client *http.Client, request SharedAlbumsLeaveRequest) error {
	outputJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", sharedAlbums.baseURL()+":leave", bytes.NewBuffer(outputJSON))
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

// SharedAlbumsLeaveRequest is a required body of the SharedAlbums.Leave method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/leave#request-body
type SharedAlbumsLeaveRequest struct {
	ShareToken string `json:"shareToken,omitempty"`
}

// - list

func (sharedAlbums sharedAlbumsRequests) List(client *http.Client, queries ...ListQuery) (SharedAlbumsListResponse, error) {
	values := url.Values{}
	for _, query := range queries {
		query(&values)
	}
	req, err := http.NewRequest("GET", sharedAlbums.baseURL(), nil)
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return SharedAlbumsListResponse{}, err
	}
	defer resp.Body.Close()
	e := RequestError(resp)
	if e != nil {
		return SharedAlbumsListResponse{}, e
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SharedAlbumsListResponse{}, err
	}
	var response SharedAlbumsListResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return SharedAlbumsListResponse{}, err
	}
	return response, nil
}

// SharedAlbumsListResponse is the body returned by the SharedAlbums.List method.
// Source: https://developers.google.com/photos/library/reference/rest/v1/sharedAlbums/list#response-body
type SharedAlbumsListResponse struct {
	SharedAlbums  []Album `json:"sharedAlbums,omitempty"`
	NextPageToken string  `json:"nextPageToken,omitempty"`
}
