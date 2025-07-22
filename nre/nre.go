package nre

import (
	"fmt"
	"ptt/output"
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

func DisplayDepartureBoard(board *nr.StationBoard) output.Table {
	table := output.Table{}
	for _, s := range board.TrainServices {
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
		if s.Platform == nil {
			platform = ""
		} else {
			platform = fmt.Sprintf("Platform %s", *s.Platform)
		}
		table.AddRow(output.NewFormattedRow(
			output.NewFormattedText(s.Destination.Name, nil),
			output.NewFormattedText(platform, nil),
			output.NewFormattedText(s.STD, nil),
			output.NewFormattedText(s.ETD, etdColor),
		))
	}
	return table
}
