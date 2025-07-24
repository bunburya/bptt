package search

import (
	"log"
	"ptt/output"
	"ptt/tfl"
	"strings"

	"github.com/spf13/cobra"
)

var SearchStopCmd = &cobra.Command{
	Use:  "stop",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		searchStr := strings.Join(args, " ")
		stopPoints, err := tfl.SearchStopPoints(searchStr)
		if err != nil {
			log.Fatal(err)
		}
		table := output.Table{}
		for _, stopPoint := range stopPoints {
			table.AddRow(stopPoint.ToRow())
		}
		table.Print("\t", true, false)
	},
}
