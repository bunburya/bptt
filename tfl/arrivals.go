package tfl

import (
	"errors"
	"fmt"
	"ptt/config"
	"ptt/output"
	"slices"
	"sort"
	"time"

	"github.com/fatih/color"
)

func arrivalsUrl(naptanId string) string {
	return fmt.Sprintf("%s/stopPoint/%s/Arrivals", BaseUrl, naptanId)
}

type arrival struct {
	StationName   string `json:"stationName"`
	LineId        string `json:"lineId"`
	LineName      string `json:"lineName"`
	Destination   string `json:"destinationName"`
	TimeToStation int    `json:"timeToStation"`
}

func (a arrival) toRow() output.Row {
	duration := time.Duration(a.TimeToStation) * time.Second
	lineCol := output.NewCell(a.LineName, nil)
	dstCol := output.NewCell(a.Destination, nil)
	timeCol := output.NewCell(fmt.Sprintf("%s", duration), nil)
	return output.NewRow(lineCol, dstCol, timeCol)
}

func filterByLine(arrivals []arrival, lines []string) []arrival {
	var filtered []arrival
	for _, a := range arrivals {
		if slices.Contains(lines, a.LineId) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func getStopArrivals(naptanId string, lines []string, count int, apiKey string) ([]arrival, error) {
	if len(naptanId) == 0 {
		return nil, errors.New("no naptanId provided")
	}
	naptanId = config.ResolveAlias("tfl.stop_point_aliases", naptanId)
	url := arrivalsUrl(naptanId)
	arrivals, err := request[[]arrival](url, apiKey)
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

func ArrivalsTable(
	naptanId string,
	lines []string,
	count int,
	apiKey string,
	options output.Options,
) (output.Table, error) {
	table := output.Table{}
	arrivals, err := getStopArrivals(naptanId, lines, count, apiKey)
	if err != nil {
		return table, err
	}
	if options.Header {
		table.AddRow(output.NewRow(
			output.NewCell("Line", color.New(color.Bold)),
			output.NewCell("Destination", color.New(color.Bold)),
			output.NewCell("ETA", color.New(color.Bold)),
		))
	}
	for _, arr := range arrivals {
		table.AddRow(arr.toRow())
	}
	if options.Timestamp {
		table.Timestamp()
	}
	return table, nil
}
