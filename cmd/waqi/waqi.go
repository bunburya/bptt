package waqi

import (
	"ptt/output"
	"ptt/waqi"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AqiCmd = &cobra.Command{
	Use:   "waqi",
	Short: "Report today's Air Quality Index for a given location.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := output.OptionsFromConfig()
		apiKey := viper.GetString("waqi.api_key")
		table, err := waqi.CityAqiTable(args[0], apiKey, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color)
		return nil
	},
}

func init() {
	AqiCmd.PersistentFlags().StringP("api-key", "k", "", "waqi.info API token")
	_ = viper.BindPFlag("waqi.api_key", AqiCmd.PersistentFlags().Lookup("api-key"))
}
