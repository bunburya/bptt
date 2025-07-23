package nre

import (
	"log"
	"os"
	"ptt/nre"

	"github.com/spf13/cobra"
)

// departuresCmd represents the departures command
var departuresCmd = &cobra.Command{
	Use:   "departures",
	Short: "View departures board for the given station",
	Long:  `View departures board for the given station. The station should be identified by its CRS code.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiToken, _ := cmd.Flags().GetString("token")
		if apiToken == "" {
			apiToken = os.Getenv("PTT_NRE_API_TOKEN")
		}
		if apiToken == "" {
			log.Fatal("National Rail API token is required")
		}
		callPoints, _ := cmd.Flags().GetStringSlice("calls")
		showPlatform, _ := cmd.Flags().GetBool("platform")
		useColor, _ := cmd.Flags().GetBool("color")
		count, _ := cmd.Flags().GetInt("count")
		depBoard, err := nre.GetDepartureBoard(args[0], apiToken)
		if err != nil {
			log.Fatal(err)
		}
		table := nre.DisplayDepartureBoard(depBoard, callPoints, showPlatform, count)
		table.Print("\t", true, useColor)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// departuresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// departuresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	departuresCmd.Flags().StringSlice("calls", nil,
		"comma-separated list of CRS codes "+
			"(only display departures for services that subsequently call at one of the specified stations)")
	departuresCmd.Flags().BoolP("platform", "p", false, "display platform number")
	departuresCmd.Flags().IntP("count", "n", 0, "max number of departures to display")
}
