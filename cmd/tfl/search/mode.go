package search

import (
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "list all mode IDs supported by the TfL API",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.ModesTable(apiKey, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color)
		return nil
	},
}
