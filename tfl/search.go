package tfl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ptt/output"
	"slices"
)

type StopPoint struct {
	Name     string   `json:"name"`
	NaptanId string   `json:"id"`
	Modes    []string `json:"modes"`
}

func stopPointSearchUrl(query string, apiKey string) string {
	query = url.PathEscape(query)
	return addApiKey(fmt.Sprintf("%s/StopPoint/Search/%s", BaseUrl, query), apiKey)
}

func filterStopsByModes(stops []StopPoint, modes []string) []StopPoint {
	var filtered []StopPoint
	for _, s := range stops {
		for _, m := range s.Modes {
			if slices.Contains(modes, m) {
				filtered = append(filtered, s)
				break
			}
		}
	}
	return filtered
}

func SearchStopPoints(query string, modes []string, apiKey string) ([]StopPoint, error) {
	searchUrl := stopPointSearchUrl(query, apiKey)
	searchResult, err := request[map[string]json.RawMessage](searchUrl, apiKey)
	if err != nil {
		return nil, err
	}
	var stopPoints []StopPoint
	if err := json.Unmarshal(searchResult["matches"], &stopPoints); err != nil {
		return nil, err
	}
	if len(modes) > 0 {
		stopPoints = filterStopsByModes(stopPoints, modes)
	}
	return stopPoints, nil
}

func (sp StopPoint) ToRow() output.Row {
	return output.NewRow(
		output.NewCell(sp.Name, nil),
		output.NewCell(sp.NaptanId, nil),
	)
}
