package tfl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ptt/output"
	"slices"
	"strings"
)

type StopPoint struct {
	Name     string   `json:"name"`
	NaptanId string   `json:"id"`
	Modes    []string `json:"modes"`
}

func stopPointSearchUrl(query string) string {
	query = url.PathEscape(query)
	return fmt.Sprintf("%s/StopPoint/Search/%s", BaseUrl, query)
}

// filterStopsByModes filters a slice of `StopPoint` structs depending on whether they serve any of the given modes.
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
	searchUrl := stopPointSearchUrl(query)
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

type mode struct {
	Name string `json:"modeName"`
}

// SearchModes fetches a list of all mode IDs from the TfL API and returns it.
func SearchModes(apiKey string) ([]string, error) {
	searchUrl := fmt.Sprintf("%s/Line/Meta/Modes", BaseUrl)
	data, err := request[[]mode](searchUrl, apiKey)
	if err != nil {
		return nil, err
	}
	var modes []string
	for _, m := range data {
		modes = append(modes, m.Name)
	}
	return modes, nil
}

func linesSearchUrl(modes []string) string {
	return fmt.Sprintf("%s/Line/Mode/%s", BaseUrl, strings.Join(modes, ","))
}

func SearchLines(modes []string, apiKey string) (map[string][]Line, error) {
	lines, err := request[[]Line](linesSearchUrl(modes), apiKey)
	if err != nil {
		return nil, err
	}
	results := make(map[string][]Line)
	for _, line := range lines {
		if results[line.Mode] == nil {
			results[line.Mode] = []Line{line}
		} else {
			results[line.Mode] = append(results[line.Mode], line)
		}
	}
	return results, nil
}

func (line *Line) ToRowWithMode() output.Row {
	return output.NewRow(
		output.NewCell(line.Name, nil),
		output.NewCell(line.Id, nil),
		output.NewCell(line.Mode, nil),
	)
}
