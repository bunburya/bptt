package nre

import (
	"errors"
	"os"
	"ptt/nre"
	"ptt/output"

	"github.com/spf13/cobra"
)

// departuresCmd represents the departures command
var departuresCmd = &cobra.Command{
	Use:   "departures",
	Short: "View departures board for the given station",
	Long:  `View departures board for the given station. The station should be identified by its CRS code.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiToken, _ := cmd.Flags().GetString("token")
		if apiToken == "" {
			apiToken = os.Getenv("PTT_NRE_API_TOKEN")
		}
		if apiToken == "" {
			return errors.New("API token is required")
		}
		callPoints, _ := cmd.Flags().GetStringSlice("calls")
		showPlatform, _ := cmd.Flags().GetBool("platform")
		count, _ := cmd.Flags().GetInt("count")
		opt := output.OptionsFromFlags(cmd.Flags())
		table, err := nre.DeparturesTable(args[0], callPoints, count, apiToken, showPlatform, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color)
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
