package tfl

import (
	"bptt/output"
	"bptt/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var modeStatusCmd = &cobra.Command{
	Use:   "modestatus",
	Short: "View the current service status of each line of the given modes",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		modeIds := args
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.ModeStatusTable(modeIds, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}

func init() {}
