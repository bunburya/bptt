package search

import (
	"log"
	"ptt/output"
	"ptt/tfl"
	"strings"

	"github.com/spf13/cobra"
)

var searchStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "search for stop points with the given string in their name",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchStr := strings.Join(args, " ")
		modes, _ := cmd.Flags().GetStringSlice("modes")
		apiKey, _ := cmd.Flags().GetString("api-key")
		stopPoints, err := tfl.SearchStopPoints(searchStr, modes, apiKey)
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

func init() {
	searchStopCmd.Flags().StringSliceP("modes", "m", nil,
		"search only for stops serving the given modes")
}
