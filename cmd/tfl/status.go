package tfl

import (
	"bptt/internal/output"
	"bptt/internal/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View the current service status of each of the given lines",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lineIds := args
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.LineStatusTable(lineIds, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}

func init() {}
