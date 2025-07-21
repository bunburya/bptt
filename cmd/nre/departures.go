package nre

import (
	"log"
	"ltt/nre"
	"os"

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
			apiToken = os.Getenv("LTT_NRE_API_TOKEN")
		}
		if apiToken == "" {
			log.Fatal("National Rail API token is required")
		}

		depBoard, err := nre.GetDepartureBoard(args[0], apiToken)
		if err != nil {
			log.Fatal(err)
		}
		table := nre.DisplayDepartureBoard(depBoard)
		table.Print("\t", true, false)
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
}
