package search

import (
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var searchModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "list all mode IDs supported by the TfL API",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		opt := output.OptionsFromFlags(cmd.Flags())
		table, err := tfl.ModesTable(apiKey, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color)
		return nil
	},
}
