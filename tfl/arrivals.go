package tfl

import (
	"encoding/json"
	"errors"
	"fmt"
	"ltt/output"
	"net/http"
	"slices"
	"sort"
	"time"
)

func arrivalsUrl(naptanId string) string {
	return fmt.Sprintf("%s/StopPoint/%s/Arrivals", BaseUrl, naptanId)
}

type Arrival struct {
	StationName   string `json:"stationName"`
	LineName      string `json:"lineName"`
	Destination   string `json:"destinationName"`
	TimeToStation int    `json:"timeToStation"`
}

func (a Arrival) ToRow() output.FormattedRow {
	duration := time.Duration(a.TimeToStation) * time.Second
	lineCol := output.NewFormattedText(a.LineName, nil)
	dstCol := output.NewFormattedText(a.Destination, nil)
	timeCol := output.NewFormattedText(fmt.Sprintf("%s", duration), nil)
	return output.NewFormattedRow(lineCol, dstCol, timeCol)
}

func filterByLine(arrivals []Arrival, lines []string) []Arrival {
	var filtered []Arrival
	for _, a := range arrivals {
		if slices.Contains(lines, a.LineName) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

type StopArrivals struct {
	NaptanId    string
	StationName string
	Arrivals    []Arrival
}

func NewStopArrivals(naptanId string, arrivals []Arrival, lines []string, count int) (*StopArrivals, error) {
	if len(arrivals) == 0 {
		return nil, errors.New("no arrivals provided")
	}
	if len(lines) > 0 {
		arrivals = filterByLine(arrivals, lines)
	}
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].TimeToStation < arrivals[j].TimeToStation
	})
	if count > 0 {
		arrivals = arrivals[:min(count, len(arrivals))]
	}
	return &StopArrivals{
		naptanId,
		arrivals[0].StationName,
		arrivals,
	}, nil
}

func GetStopArrivals(naptanId string, lines []string, count int) (*StopArrivals, error) {
	url := arrivalsUrl(naptanId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ltt")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var arrivals []Arrival
	if err := json.NewDecoder(resp.Body).Decode(&arrivals); err != nil {
		return nil, err
	}
	return NewStopArrivals(naptanId, arrivals, lines, count)
}
