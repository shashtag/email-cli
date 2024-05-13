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

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "soc-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func addPalette() {
	// Add the command to the root command
	rootCmd.AddCommand(createemail.CreateEmailCmd)

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	if err := viper.WriteConfigAs("./config.yaml"); err != nil {
		fmt.Println(err)
	}
	addPalette()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

type Config struct {
	SMTP_SERVER string `mapstructure:"SMTP_SERVER"`
}

func initConfig() {

	viper.SetConfigFile("../app.env") // set the path of your environment file

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// viper.Set("ddmdm", "ss")
	// if err := viper.WriteConfigAs("../app.env"); err != nil {
	// 	fmt.Println(err)
	// }

}
