package tfl

import (
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var bikesCmd = &cobra.Command{
	Use:   "bikes",
	Short: "display status (number of available bikes and empty slots) at the given bike stations",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		opt := output.OptionsFromFlags(cmd.Flags())
		t, err := tfl.BikesTable(args, apiKey, opt)
		if err != nil {
			return err
		}
		t.Print("\t", true, opt.Color)
		return nil
	},
}
