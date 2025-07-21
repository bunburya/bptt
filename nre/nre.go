package nre

import (
	"ltt/output"

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
		table.AddRow(output.NewFormattedRow(
			output.NewFormattedText(s.Destination.Name, nil),
			output.NewFormattedText(s.STD, nil),
			output.NewFormattedText(s.ETD, nil),
		))
	}
	return table
}
