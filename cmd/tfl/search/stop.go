package search

import (
	"bptt/output"
	"bptt/tfl"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "search for stop points with the given string in their name",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		searchStr := strings.Join(args, " ")
		modes, _ := cmd.Flags().GetStringSlice("modes")
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.StopPointsTable(searchStr, modes, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}

func init() {
	searchStopCmd.Flags().StringSliceP("modes", "m", nil,
		"search only for stops serving the given modes")
}
