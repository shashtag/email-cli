/*
Copyright © 2024 Shashwat Gupta shashwatsatna@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	createemail "github.com/shashtag/soc-cli/cmd/createEmail"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "soc-cli",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func addPalette() {
	// Add all the palettes of commands
	rootCmd.AddCommand(createemail.CreateEmailCmd)

}

func init() {
	cobra.OnInitialize(initConfig)
	addPalette()
}

func initConfig() {

	viper.SetConfigFile("../app.env")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// viper.Set("ddmdm", "ss")
	// if err := viper.WriteConfigAs("../app.env"); err != nil {
	// 	fmt.Println(err)
	// }

}
