package search

import "github.com/spf13/cobra"

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search various TfL resources",
}

func init() {
	SearchCmd.AddCommand(searchStopCmd)
	SearchCmd.AddCommand(searchModeCmd)
	SearchCmd.AddCommand(searchLineCmd)
	SearchCmd.AddCommand(searchBikeCmd)
}
