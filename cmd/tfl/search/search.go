package search

import "github.com/spf13/cobra"

var SearchCmd = &cobra.Command{
	Use: "search",
}

func init() {
	SearchCmd.AddCommand(searchStopCmd)
	SearchCmd.AddCommand(searchModeCmd)
	SearchCmd.AddCommand(searchLineCmd)
	SearchCmd.AddCommand(searchBikeCmd)
}
