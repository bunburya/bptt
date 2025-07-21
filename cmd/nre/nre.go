/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package nre

import (
	"github.com/spf13/cobra"
)

// NreCmd represents the nre command
var NreCmd = &cobra.Command{
	Use:   "nre",
	Short: "Access information about National Rail services",
}

func init() {
	NreCmd.AddCommand(departuresCmd)

	// Here you will define your flags and configuration settings.
	NreCmd.PersistentFlags().StringP("token", "t", "", "National Rail API token")

}
