package search

import (
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var searchLineCmd = &cobra.Command{
	Use:   "line",
	Short: "search for all lines of the given modes",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		opt := output.OptionsFromFlags(cmd.Flags())
		table, err := tfl.LinesTable(args, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color)
		return nil
	},
}
