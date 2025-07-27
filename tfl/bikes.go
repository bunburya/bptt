package tfl

import (
	"encoding/json"
	"fmt"
	"ptt/output"
	"slices"
	"strconv"

	"github.com/fatih/color"
)

const bikePointUrl = BaseUrl + "/bikePoint"

func singleBikePointUrl(id string) string {
	return fmt.Sprintf("%s/%s", bikePointUrl, id)
}

type tmpBikePoint struct {
	Id         string              `json:"id"`
	Name       string              `json:"commonName"`
	Properties []map[string]string `json:"additionalProperties"`
}

type bikePoint struct {
	Id            string `json:"id"`
	Name          string `json:"commonName"`
	Ebikes        int
	StandardBikes int
	EmptyDocks    int
}

func (b *bikePoint) UnmarshalJSON(data []byte) error {
	var tmp tmpBikePoint
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	b.Id = tmp.Id
	b.Name = tmp.Name
	for _, p := range tmp.Properties {
		switch p["key"] {
		case "NbEmptyDocks":
			b.EmptyDocks, err = strconv.Atoi(p["value"])
			if err != nil {
				return err
			}
		case "NbStandardBikes":
			b.StandardBikes, err = strconv.Atoi(p["value"])
			if err != nil {
				return err
			}
		case "NbEBikes":
			b.Ebikes, err = strconv.Atoi(p["value"])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *bikePoint) toRow() output.Row {
	return output.NewRow(
		output.NewCell(b.Name, nil),
		output.NewCell(strconv.Itoa(b.StandardBikes), nil),
		output.NewCell(strconv.Itoa(b.Ebikes), nil),
		output.NewCell(strconv.Itoa(b.EmptyDocks), nil),
	)
}

func getSingleBikePoint(id string, apiKey string) (bikePoint, error) {
	url := singleBikePointUrl(id)
	return request[bikePoint](url, apiKey)
}

func filterByBikePointId(bp []bikePoint, ids []string) []bikePoint {
	var filtered []bikePoint
	for _, b := range bp {
		if slices.Contains(ids, b.Id) {
			filtered = append(filtered, b)
		}
	}
	return filtered
}

func getMultiBikePoints(ids []string, apiKey string) ([]bikePoint, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	bps, err := request[[]bikePoint](bikePointUrl, apiKey)
	if err != nil {
		return nil, err
	}
	bps = filterByBikePointId(bps, ids)
	return bps, nil
}

func getBikePoints(bpIds []string, apiKey string) ([]bikePoint, error) {
	if len(bpIds) == 1 {
		bp, err := getSingleBikePoint(bpIds[0], apiKey)
		if err != nil {
			return nil, err
		}
		return []bikePoint{bp}, nil
	} else {
		return getMultiBikePoints(bpIds, apiKey)
	}
}

func BikesTable(bikePointIds []string, apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	bikePoints, err := getBikePoints(bikePointIds, apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.AddRow(output.NewRow(
			output.NewCell("Station", color.New(color.Bold)),
			output.NewCell("Bikes", color.New(color.Bold)),
			output.NewCell("E-Bikes", color.New(color.Bold)),
			output.NewCell("Empty docks", color.New(color.Bold)),
		))
	}
	for _, bp := range bikePoints {
		t.AddRow(bp.toRow())
	}
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}
