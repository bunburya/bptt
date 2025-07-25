package tfl

import (
	"log"
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View the current service status of each of the given lines",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lineIds := args
		useColor, _ := cmd.Flags().GetBool("color")
		apiKey, _ := cmd.Flags().GetString("api-key")
		lines, err := tfl.GetLineStatuses(lineIds, apiKey)
		if err != nil {
			log.Fatal(err)
		}
		table := output.Table{}
		for _, line := range lines {
			row, err := line.ToRowWithStatus(useColor)
			if err != nil {
				log.Fatal(err)
			}
			table.AddRow(row)
		}
		table.Print("\t", true, useColor)
	},
}

func init() {}
