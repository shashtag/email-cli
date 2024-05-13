/*
Copyright © 2024 Shashwat Gupta shashwatsatna@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
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

	// viper.Set("aa", "smtp.gmail.com")
	fmt.Println(viper.Get("aa"))
	addPalette()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Config struct {
	SMTP_SERVER string `mapstructure:"SMTP_SERVER"`
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Search config in parent directory with name "app.env"
	viper.AddConfigPath("..")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv() // read in environment variables that match

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

}
