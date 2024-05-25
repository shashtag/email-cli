/*
Copyright © 2024 Shashwat Gupta shashwatsatna@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/shashtag/email-cli/cmd/createEmail/wsgm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "email-cli",
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

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(wsgm.WsgmCmd)
}

func initConfig() {

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetConfigFile(home + "/email-cli.env")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	setRootDirectory()
}

func setRootDirectory() {
	//set the project root directory in viper
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../")
	viper.Set("PROJECT_ROOT_DIR", d)
}
