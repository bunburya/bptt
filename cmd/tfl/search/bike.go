package search

import (
	"bptt/output"
	"bptt/tfl"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchBikeCmd = &cobra.Command{
	Use:   "bike",
	Short: "search for bike points with the given string in their name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		searchStr := strings.Join(args, " ")
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.BikePointsTable(searchStr, apiKey, opt)
		if err != nil {
			return nil
		}
		table.Print(opt)
		return nil
	},
}
