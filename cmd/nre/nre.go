package nre

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NreCmd represents the nre command
var NreCmd = &cobra.Command{
	Use:   "nre",
	Short: "Access information about National Rail services",
}

func init() {
	NreCmd.AddCommand(departuresCmd)

	NreCmd.PersistentFlags().StringP("api-key", "k", "", "National Rail API token")
	_ = viper.BindPFlag("nre.api_key", NreCmd.PersistentFlags().Lookup("api-key"))

}
