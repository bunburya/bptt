package tfl

import (
	"bptt/internal/output"
	"bptt/internal/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var bikesCmd = &cobra.Command{
	Use:   "bikes",
	Short: "Display status (number of available bikes and empty slots) at the given bike stations",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.BikesTable(args, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}
