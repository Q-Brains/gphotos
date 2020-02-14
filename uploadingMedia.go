package gphotos

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

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

func (upload uploadingMediaRequests) UploadMedia(client *http.Client, filePath string, filename string) (uploadToken string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", upload.baseURL(), file)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("X-Goog-Upload-File-Name", filename)
	req.Header.Set("X-Goog-Upload-Protocol", "raw")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return bytes.NewBuffer(b).String(), nil
}

func (upload uploadingMediaRequests) ResumableUploads(client *http.Client, filePath string, filename string) (uploadToken string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", upload.baseURL(), nil)
	if err != nil {
		return "", err
	}
	contentType, err := detectContentType(file)
	if err != nil {
		return "", err
	}
	length, err := byteLength(file)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Length", strconv.Itoa(0))
	req.Header.Set("X-Goog-Upload-Command", "start")
	req.Header.Set("X-Goog-Upload-Content-Type", contentType)
	req.Header.Set("X-Goog-Upload-File-Name", filename)
	req.Header.Set("X-Goog-Upload-Protocol", "resumable")
	req.Header.Set("X-Goog-Upload-Raw-Size", strconv.FormatInt(length, 10))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	uploadURL := resp.Header.Get("X-Goog-Upload-URL")
	req, err = http.NewRequest("POST", uploadURL, file)
	req.Header.Set("Content-Length", strconv.FormatInt(length, 10))
	req.Header.Set("X-Goog-Upload-Command", "upload, finalize")
	req.Header.Set("X-Goog-Upload-Offset", strconv.Itoa(0))
	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return "", nil
}

func detectContentType(file *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buffer), nil
}

func byteLength(file *os.File) (int64, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
