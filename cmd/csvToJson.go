/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// csvToJsonCmd represents the csvToJson command
var csvToJsonCmd = &cobra.Command{
	Use:   "csvToJson",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("csvToJson called")
		csvToJson(cmd, args)
	},
}

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func init() {
	rootCmd.AddCommand(csvToJsonCmd)

	// Here you will define your flags and configuration settings.
	csvToJsonCmd.Flags().String("separator", "comma", "Column separator")
	csvToJsonCmd.Flags().Bool("pretty", false, "Generate pretty JSON")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csvToJsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvToJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func csvToJson(cmd *cobra.Command, args []string) {
	separator, _ := cmd.Flags().GetString("separator")
	pretty, _ := cmd.Flags().GetBool("pretty")
	makeInputFile(args, separator, pretty)
}
func makeInputFile(args []string, separator string, pretty bool) {
	if !(separator == "comma" || separator == "semicolon") {
		log.Fatal("Only comma or semicolon separators are allowed")
	}
	fmt.Println(args)
}
func checkIfValidFile(filename string) (bool, error) {
	// Checking if entered file is CSV by using the filepath package from the standard library
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("File %s is not CSV", filename)
	}

	// Checking if filepath entered belongs to an existing file. We use the Stat method from the os package (standard library)
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist", filename)
	}
	// If we get to this point, it means this is a valid file
	return true, nil
}
