/*
Copyright © 2025 Nicholas Pazienza
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CoinMarketCap-Cryptocurrency-CLI",
	Short: "Allows users to retrieve data from CoinMarketCap.",
	Long:  `User can enter desired Cryptocurrency name. Application leverages CoinMarketCap API and displays Cryptocurrency name, Symbol and current USD price. `,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
