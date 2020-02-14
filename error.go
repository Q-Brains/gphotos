package gphotos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrorResponse represents an error condition.
type ErrorResponse struct {
	Error struct {
		Code    json.Number `json:"code,omitempty"`
		Message string      `json:"message,omitempty"`
		Status  string      `json:"status,omitempty"`
	} `json:"error,omitempty"`
}

// RequestError receives an error and displays an ErrorResponse.
func RequestError(resp *http.Response) error {
	if resp.StatusCode/100 != 2 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var response ErrorResponse
		if err := json.Unmarshal(b, &response); err != nil {
			return err
		}
		fmt.Println(response)
	}
	return nil
}
