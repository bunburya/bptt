package tfl

import (
	"ptt/cmd/tfl/search"

	"github.com/spf13/cobra"
)

// TflCmd represents the tfl command
var TflCmd = &cobra.Command{
	Use:   "tfl",
	Short: "Access information about Transport for London (TfL) services",
}

func init() {
	TflCmd.AddCommand(statusCmd)
	TflCmd.AddCommand(arrivalsCmd)
	TflCmd.AddCommand(search.SearchCmd)
	TflCmd.PersistentFlags().StringP("api-key", "k", "", "TfL API key")
}
