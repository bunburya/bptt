package search

import (
	"fmt"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

var searchModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "list all mode IDs supported by the TfL API",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		modes, err := tfl.SearchModes(apiKey)
		if err != nil {
			return err
		}
		for _, m := range modes {
			fmt.Println(m)
		}
		return nil
	},
}
