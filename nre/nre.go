package nre

import (
	"bptt/config"
	"bptt/output"
	"fmt"
	"slices"
	"strings"

	"github.com/fatih/color"
	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func getDepartureBoard(crs string, apiKey string) (*nr.StationBoard, error) {
	client, err := nr.NewClient(nr.AccessTokenOpt(apiKey))
	if err != nil {
		return nil, err
	}
	return client.GetDeparturesWithDetails(nr.CRSCode(crs))
}

func resolveCallPointAliases(crs []string) []string {
	var resolved []string
	for _, c := range crs {
		resolved = append(resolved, config.ResolveAlias("nre.station_aliases", c))
	}
	return resolved
}

func filterServicesByCallPoint(services []*nr.TrainService, callPoints []string) []*nr.TrainService {
	var filtered []*nr.TrainService
	callPoints = resolveCallPointAliases(callPoints)
	for _, service := range services {
		if service == nil {
			continue
		}
		for _, callPoint := range service.SubsequentCallingPoints {
			if slices.Contains(callPoints, callPoint.CRS) {
				filtered = append(filtered, service)
				break
			}
		}
	}
	return filtered
}

func DeparturesTable(
	crs string,
	callPoints []string,
	count int,
	apiKey string,
	showPlatform bool,
	options output.Options,
) (output.Table, error) {
	table := output.Table{}
	crs = config.ResolveAlias("nre.station_aliases", crs)
	board, err := getDepartureBoard(crs, apiKey)
	if err != nil {
		return table, err
	}

	services := board.TrainServices
	if len(callPoints) > 0 {
		services = filterServicesByCallPoint(services, callPoints)
	}
	if count > 0 {
		services = services[:min(count, len(services))]
	}

	if options.Header {
		var pHeader string
		if showPlatform {
			pHeader = "Platform"
		} else {
			pHeader = ""
		}
		table.SetHeader(output.NewRow(
			output.NewCell("Destination", color.New(color.Bold)),
			output.NewCell(pHeader, color.New(color.Bold)),
			output.NewCell("STD", color.New(color.Bold)),
			output.NewCell("ETD", color.New(color.Bold)),
		))
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
	if options.Timestamp {
		table.Timestamp()
	}
	return table, nil
}
