package search

import (
	"fmt"
	"log"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var searchModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "list all mode IDs supported by the TfL API",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, _ := cmd.Flags().GetString("api-key")
		modes, err := tfl.SearchModes(apiKey)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range modes {
			fmt.Println(m)
		}
	},
}
