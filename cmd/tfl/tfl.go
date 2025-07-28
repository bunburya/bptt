package tfl

import (
	"ptt/cmd/tfl/search"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TflCmd represents the tfl command
var TflCmd = &cobra.Command{
	Use:   "tfl",
	Short: "Access information about Transport for London (TfL) services",
}

func init() {
	TflCmd.AddCommand(statusCmd)
	TflCmd.AddCommand(arrivalsCmd)
	TflCmd.AddCommand(bikesCmd)
	TflCmd.AddCommand(search.SearchCmd)
	TflCmd.PersistentFlags().StringP("api-key", "k", "", "TfL API key")
	_ = viper.BindPFlag("tfl.api_key", TflCmd.PersistentFlags().Lookup("api-key"))
}
