package tfl

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BaseUrl = "https://api.tfl.gov.uk"

// addApiKey adds the given Tfl API key to a request URL, if an API key is present. Otherwise, it returns the given
// URL unchanged. Assumes `url` is a valid TfL request URL.
func addApiKey(url string, apiKey string) string {
	if len(apiKey) > 0 {
		url += "?app_key=" + apiKey
	}
	return url
}

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
