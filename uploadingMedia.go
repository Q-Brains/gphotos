package gphotos

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

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
