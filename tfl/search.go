package tfl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"ptt/output"
)

type StopPoint struct {
	Name     string `json:"name"`
	NaptanId string `json:"id"`
}

func stopPointSearchUrl(query string) string {
	query = url.PathEscape(query)
	return fmt.Sprintf("%s/StopPoint/Search/%s", BaseUrl, query)
}

func SearchStopPoints(query string) ([]StopPoint, error) {
	searchUrl := stopPointSearchUrl(query)
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ptt")
	//println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)
	//println(string(body))

	var searchResult map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}

	var stopPoints []StopPoint
	if err := json.Unmarshal(searchResult["matches"], &stopPoints); err != nil {
		return nil, err
	}
	return stopPoints, nil
}

func (sp StopPoint) ToRow() output.Row {
	return output.NewRow(
		output.NewCell(sp.Name, nil),
		output.NewCell(sp.NaptanId, nil),
	)
}
