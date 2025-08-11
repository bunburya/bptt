package waqi

import (
	"bptt/output"
	"bptt/waqi"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AqiCmd = &cobra.Command{
	Use:   "waqi",
	Short: "Report today's Air Quality Index for a given location.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := output.OptionsFromConfig()
		apiKey := viper.GetString("waqi.api_key")
		table, err := waqi.CityAqiTable(strings.Join(args, " "), apiKey, opt)
		if err != nil {
			return err
		}
		table.Print(opt)
		return nil
	},
}

func init() {
	AqiCmd.PersistentFlags().StringP("api-key", "k", "", "waqi.info API token")
	_ = viper.BindPFlag("waqi.api_key", AqiCmd.PersistentFlags().Lookup("api-key"))
}
