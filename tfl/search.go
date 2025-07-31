package tfl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ptt/output"
	"slices"
	"strings"

	"github.com/fatih/color"
)

func stopPointSearchUrl(query string) string {
	query = url.PathEscape(query)
	return fmt.Sprintf("%s/stopPoint/Search/%s", BaseUrl, query)
}

type stopPoint struct {
	Name     string   `json:"name"`
	NaptanId string   `json:"id"`
	Modes    []string `json:"modes"`
}

func (sp *stopPoint) ToRow() output.Row {
	return output.NewRow(
		output.NewCell(sp.Name, nil),
		output.NewCell(sp.NaptanId, nil),
	)
}

// filterStopsByModes filters a slice of `stopPoint` structs depending on whether they serve any of the given modes.
func filterStopsByModes(stops []stopPoint, modes []string) []stopPoint {
	var filtered []stopPoint
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

func searchStopPoints(query string, modes []string, apiKey string) ([]stopPoint, error) {
	searchUrl := stopPointSearchUrl(query)
	searchResult, err := request[map[string]json.RawMessage](searchUrl, apiKey)
	if err != nil {
		return nil, err
	}
	var stopPoints []stopPoint
	if err := json.Unmarshal(searchResult["matches"], &stopPoints); err != nil {
		return nil, err
	}
	if len(modes) > 0 {
		stopPoints = filterStopsByModes(stopPoints, modes)
	}
	return stopPoints, nil
}

func StopPointsTable(query string, modes []string, apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	stopPoints, err := searchStopPoints(query, modes, apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.AddRow(output.NewRow(
			output.NewCell("Name", color.New(color.Bold)),
			output.NewCell("NaPTAN ID", color.New(color.Bold)),
		))
	}
	for _, s := range stopPoints {
		t.AddRow(s.ToRow())
	}
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}

type mode struct {
	Name string `json:"modeName"`
}

// searchModes fetches a list of all mode IDs from the TfL API and returns it.
func searchModes(apiKey string) ([]string, error) {
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

func ModesTable(apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	modes, err := searchModes(apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.AddRow(output.NewRow(
			output.NewCell("Mode ID", color.New(color.Bold)),
		))
	}
	for _, m := range modes {
		t.AddRow(output.NewRow(output.NewCell(m, nil)))
	}
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}

func linesSearchUrl(modes []string) string {
	return fmt.Sprintf("%s/Line/Mode/%s", BaseUrl, strings.Join(modes, ","))
}

type Line struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Mode string `json:"modeName"`
}

func (line *Line) ToRow() output.Row {
	return output.NewRow(
		output.NewCell(line.Name, nil),
		output.NewCell(line.Id, nil),
		output.NewCell(line.Mode, nil),
	)
}

func searchLines(modes []string, apiKey string) (map[string][]Line, error) {
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

func LinesTable(modes []string, apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	results, err := searchLines(modes, apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.AddRow(output.NewRow(
			output.NewCell("Name", color.New(color.Bold)),
			output.NewCell("ID", color.New(color.Bold)),
			output.NewCell("Mode", color.New(color.Bold)),
		))
	}
	for _, lines := range results {
		for _, line := range lines {
			t.AddRow(line.ToRow())
		}
	}
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}

func bikePointSearchUrl(query string) string {
	query = url.QueryEscape(query)
	return fmt.Sprintf("%s/Search?query=%s", bikePointUrl, query)
}

type bikePoint struct {
	Id   string  `json:"id"`
	Name string  `json:"commonName"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

func (point *bikePoint) ToRow() output.Row {
	return output.NewRow(
		output.NewCell(point.Name, nil),
		output.NewCell(point.Id, nil),
		output.NewCell(fmt.Sprintf("%f", point.Lat), nil),
		output.NewCell(fmt.Sprintf("%f", point.Lon), nil),
	)
}

func searchBikePoints(query string, apiKey string) ([]bikePoint, error) {
	return request[[]bikePoint](bikePointSearchUrl(query), apiKey)
}

func BikePointsTable(query string, apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	bps, err := searchBikePoints(query, apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.SetHeader(output.NewRow(
			output.NewCell("Name", color.New(color.Bold)),
			output.NewCell("ID", color.New(color.Bold)),
			output.NewCell("Latitude", color.New(color.Bold)),
			output.NewCell("Longitude", color.New(color.Bold)),
		))
	}
	for _, bp := range bps {
		t.AddRow(bp.ToRow())
	}
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}
