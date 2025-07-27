package search

import (
	"ptt/output"
	"ptt/tfl"
	"strings"

	"github.com/spf13/cobra"
)

var searchBikeCmd = &cobra.Command{
	Use:   "bike",
	Short: "search for bike points with the given string in their name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		searchStr := strings.Join(args, " ")
		apiKey, _ := cmd.Flags().GetString("api-key")
		opt := output.OptionsFromFlags(cmd.Flags())
		table, err := tfl.BikePointsTable(searchStr, apiKey, opt)
		if err != nil {
			return nil
		}
		table.Print("\t", true, opt.Color)
		return nil
	},
}
