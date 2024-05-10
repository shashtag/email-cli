/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package createemail

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CreateEmailCmd represents the createEmail command
var CreateEmailCmd = &cobra.Command{
	Use:   "createEmail",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createEmail called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createEmailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createEmailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
