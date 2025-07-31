package tfl

import (
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// arrivalsCmd represents the arrivals command
var arrivalsCmd = &cobra.Command{
	Use:   "arrivals",
	Short: "Display next arrivals at the given stop",
	Long:  "Display next arrivals at the given stop. The stop must be identified by its NaPTAN ID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lines, _ := cmd.Flags().GetStringSlice("lines")
		count, _ := cmd.Flags().GetInt("count")
		apiKey := viper.GetString("tfl.api_key")
		opt := output.OptionsFromConfig()
		table, err := tfl.ArrivalsTable(args[0], lines, count, apiKey, opt)
		if err != nil {
			return err
		}
		table.Print("\t", true, opt.Color, "no data available")
		return nil
	},
}

func init() {
	arrivalsCmd.Flags().IntP("count", "n", 0, "max number of arrivals to display")
	arrivalsCmd.Flags().StringSlice("lines", nil,
		"comma-delimited list of lines/routes to display (if not provided, all lines will be displayed)")
}
