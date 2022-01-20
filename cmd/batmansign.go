/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"giali.com/commander/cmd/template"
	"giali.com/commander/cmd/util"
	"github.com/spf13/cobra"
)

// batmansignCmd represents the batmansign command
var batmansignCmd = &cobra.Command{
	Use:   "batmansign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		template.PrintLogo(handleUserName(args))
		menu_res := template.PrintMenuAndGetResult([]string{
			"exit",
			"back",
			"continue",
		})
		handleMenuSelect(menu_res)
	},
}

func init() {
	rootCmd.AddCommand(batmansignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batmansignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batmansignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func handleUserName(args []string) string {
	if len(args) > 0 {
		return args[0]
	} else {
		return "of code"
	}
}
func handleMenuSelect(eventValue string) string {
	value := strings.TrimRight(eventValue, "\n")
	if value == "" {
		return util.CONTINUE
	} else if value == "continue" {
		return util.CONTINUE
	} else if value == "back" {
		return util.BACK
	} else if value == "exit" {
		return util.EXIT
	} else {
		fmt.Println("WARNING WRONG INPUT")
		return util.CONTINUE
	}
}
