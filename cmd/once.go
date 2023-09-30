/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// onceCmd represents the once command
var onceCmd = &cobra.Command{
	Use:   "once",
	Short: "Send a GET request once to each URL specified in configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("once called")
	},
}

func init() {
	rootCmd.AddCommand(onceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
