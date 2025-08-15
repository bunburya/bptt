package nre

import (
	"bptt/internal/nre"
	"bptt/internal/output"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// departuresCmd represents the departures command
var departuresCmd = &cobra.Command{
	Use:   "departures",
	Short: "View departures board for the given station",
	Long:  `View departures board for the given station. The station should be identified by its CRS code.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		//apiToken, _ := cmd.Flags().GetString("api-key")
		apiToken := viper.GetString("nre.api_key")
		if apiToken == "" {
			return errors.New("API key is required")
		}
		callPoints, _ := cmd.Flags().GetStringSlice("calls")
		count, _ := cmd.Flags().GetInt("count")
		opt := output.OptionsFromConfig()
		table, err := nre.DeparturesTable(args[0], callPoints, count, apiToken, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}

func init() {
	departuresCmd.Flags().StringSlice("calls", nil,
		"comma-separated list of CRS codes "+
			"(only display departures for services that subsequently call at one of the specified stations)")
	departuresCmd.Flags().BoolP("platform", "p", false, "display platform number")
	departuresCmd.Flags().IntP("count", "n", 0, "max number of departures to display")
}
