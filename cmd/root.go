package cmd

import (
	"os"
	"ptt/cmd/nre"
	"ptt/cmd/tfl"
	"ptt/cmd/waqi"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ptt",
	Short: "Access information about public transport in the terminal.",
	Long: `ptt is a command-line tool to easily access information about the current status of various public transport
services. Currently the focus is on services available in and around London, UK. For example, you can view the current
service status of tube lines, the next arrivals at your local bus stop or departures from your local train station.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(tfl.TflCmd)
	rootCmd.AddCommand(nre.NreCmd)
	rootCmd.AddCommand(waqi.AqiCmd)

	rootCmd.PersistentFlags().Bool("color", false, "use colour in output (where possible)")
	_ = viper.BindPFlag("color", rootCmd.PersistentFlags().Lookup("color"))
	rootCmd.PersistentFlags().Bool("header", false, "print header row")
	_ = viper.BindPFlag("header", rootCmd.PersistentFlags().Lookup("header"))
	rootCmd.PersistentFlags().Bool("timestamp", false, "print last updated time as footer")
	_ = viper.BindPFlag("timestamp", rootCmd.PersistentFlags().Lookup("timestamp"))

}
