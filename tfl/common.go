package tfl

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BaseUrl = "https://api.tfl.gov.uk"

// request makes a single request of the TfL API. `url` must be a valid API endpoint and `apiKey` must either be an
// empty string or a valid API key. The type parameter specifies the type of object that will be returned. It must be
// capable of being unmarshalled from the JSON response.
func request[T any](url string, apiKey string) (T, error) {
	var t T
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return t, err
	}
	req.Header.Set("User-Agent", "ptt")
	if len(apiKey) > 0 {
		req.Header.Set("app_key", apiKey)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return t, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return t, errors.New(resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

// Line represents a single TfL line/route. It is used by both the `status` and `search line` commands.
type Line struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Mode     string        `json:"modeName"`
	Statuses []*LineStatus `json:"lineStatuses"`
}
