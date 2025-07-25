package tfl

import (
	"errors"
	"fmt"
	"ptt/output"
	"slices"
	"sort"
	"time"
)

func arrivalsUrl(naptanId string, apiKey string) string {
	return addApiKey(fmt.Sprintf("%s/StopPoint/%s/Arrivals", BaseUrl, naptanId), apiKey)
}

type Arrival struct {
	StationName   string `json:"stationName"`
	LineId        string `json:"lineId"`
	LineName      string `json:"lineName"`
	Destination   string `json:"destinationName"`
	TimeToStation int    `json:"timeToStation"`
}

func (a Arrival) ToRow() output.Row {
	duration := time.Duration(a.TimeToStation) * time.Second
	lineCol := output.NewCell(a.LineName, nil)
	dstCol := output.NewCell(a.Destination, nil)
	timeCol := output.NewCell(fmt.Sprintf("%s", duration), nil)
	return output.NewRow(lineCol, dstCol, timeCol)
}

func filterByLine(arrivals []Arrival, lines []string) []Arrival {
	var filtered []Arrival
	for _, a := range arrivals {
		if slices.Contains(lines, a.LineId) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func GetStopArrivals(naptanId string, lines []string, count int, apiKey string) ([]Arrival, error) {
	if len(naptanId) == 0 {
		return nil, errors.New("no naptanId provided")
	}
	url := arrivalsUrl(naptanId, apiKey)
	arrivals, err := request[[]Arrival](url, apiKey)
	if err != nil {
		return nil, err
	}

	// Filter to relevant lines if they were provided
	if len(lines) > 0 {
		arrivals = filterByLine(arrivals, lines)
	}
	// Sort by ETA
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].TimeToStation < arrivals[j].TimeToStation
	})
	// Limit to desired number
	if count > 0 {
		arrivals = arrivals[:min(count, len(arrivals))]
	}
	return arrivals, nil
}
