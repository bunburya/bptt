package nre

import (
	"fmt"
	"ptt/output"
	"slices"
	"strings"

	"github.com/fatih/color"
	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func GetDepartureBoard(crs string, apiKey string) (*nr.StationBoard, error) {
	client, err := nr.NewClient(nr.AccessTokenOpt(apiKey))
	if err != nil {
		return nil, err
	}
	return client.GetDeparturesWithDetails(nr.CRSCode(crs))
}

func filterServicesByCallPoint(services []*nr.TrainService, dests []string) []*nr.TrainService {
	var filtered []*nr.TrainService
	for _, service := range services {
		if service == nil {
			continue
		}
		for _, callPoint := range service.SubsequentCallingPoints {
			if slices.Contains(dests, callPoint.CRS) {
				filtered = append(filtered, service)
				break
			}
		}
	}
	return filtered
}

func DisplayDepartureBoard(
	board *nr.StationBoard,
	callPoints []string,
	showPlatform bool,
	count int,
) output.Table {

	table := output.Table{}
	services := board.TrainServices
	if len(callPoints) > 0 {
		services = filterServicesByCallPoint(services, callPoints)
	}
	if count > 0 {
		services = services[:min(count, len(services))]
	}
	for _, s := range services {
		// TODO: Double check that this is sufficient to determine if a service is delayed or cancelled.
		// https://wiki.openraildata.com/index.php/GetDepBoardWithDetails suggests that delayed trains may not display
		// an ETD.
		var etdColor *color.Color
		lowerEtd := strings.ToLower(s.ETD)
		if (lowerEtd == "on time") || s.ETD == s.STD {
			etdColor = output.SafetyColors["green"]
		} else if lowerEtd == "cancelled" {
			etdColor = output.SafetyColors["red"]
		} else {
			// Apparently not on time or cancelled, so presumably delayed
			etdColor = output.SafetyColors["yellow"]
		}
		if etdColor != nil {
			etdColor = etdColor.Add(color.Bold)
		}
		var platform string
		if (!showPlatform) || (s.Platform == nil) {
			platform = ""
		} else {
			platform = fmt.Sprintf("Platform %s", *s.Platform)
		}
		table.AddRow(output.NewRow(
			output.NewCell(s.Destination.Name, nil),
			output.NewCell(platform, nil),
			output.NewCell(s.STD, nil),
			output.NewCell(s.ETD, etdColor),
		))
	}
	return table
}
