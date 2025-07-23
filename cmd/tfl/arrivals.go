/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package tfl

import (
	"log"
	"ptt/output"
	"ptt/tfl"

	"github.com/spf13/cobra"
)

// arrivalsCmd represents the arrivals command
var arrivalsCmd = &cobra.Command{
	Use:   "arrivals",
	Short: "Display next arrivals at the given stop",
	Long:  "Display next arrivals at the given stop. The stop must be identified by its NaPTAN ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lines, _ := cmd.Flags().GetStringSlice("lines")
		count, _ := cmd.Flags().GetInt("count")
		arrivals, err := tfl.GetStopArrivals(args[0], lines, count)
		if err != nil {
			log.Fatal(err)
		}

		table := output.Table{}
		for _, arr := range arrivals {
			row := arr.ToRow()
			table.AddRow(row)
		}
		table.Print("\t", true, false)
	},
}

func init() {
	arrivalsCmd.Flags().IntP("count", "n", 0, "max number of arrivals to display")
	arrivalsCmd.Flags().StringSlice("lines", nil,
		"comma-delimited list of lines/routes to display (if not provided, all lines will be displayed)")
}
