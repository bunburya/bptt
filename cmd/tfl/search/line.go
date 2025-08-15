package search

import (
	"bptt/internal/output"
	"bptt/internal/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchLineCmd = &cobra.Command{
	Use:   "line",
	Short: "search for all lines of the given modes",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.LinesTable(args, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}
