package search

import (
	"log"
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var searchLineCmd = &cobra.Command{
	Use:   "line",
	Short: "search for all lines of the given modes",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, _ := cmd.Flags().GetString("api-key")
		results, err := tfl.SearchLines(args, apiKey)
		if err != nil {
			log.Fatal(err)
		}
		table := output.Table{}
		for _, mode := range results {
			for _, line := range mode {
				table.AddRow(line.ToRowWithMode())
			}
		}
		table.Print("\t", true, false)
	},
}
