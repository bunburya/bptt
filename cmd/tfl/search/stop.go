package search

import (
	"ptt/output"
	"ptt/tfl"
	"strings"

	"github.com/spf13/cobra"
)

var searchStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "search for stop points with the given string in their name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		searchStr := strings.Join(args, " ")
		modes, _ := cmd.Flags().GetStringSlice("modes")
		apiKey, _ := cmd.Flags().GetString("api-key")
		stopPoints, err := tfl.SearchStopPoints(searchStr, modes, apiKey)
		if err != nil {
			return err
		}
		table := output.Table{}
		for _, stopPoint := range stopPoints {
			table.AddRow(stopPoint.ToRow())
		}
		table.Print("\t", true, false)
		return nil
	},
}

func init() {
	searchStopCmd.Flags().StringSliceP("modes", "m", nil,
		"search only for stops serving the given modes")
}
