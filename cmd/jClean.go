/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"giali.com/commander/cmd/util"
	"github.com/spf13/cobra"
)

// jCleanCmd represents the jClean command
var jCleanCmd = &cobra.Command{
	Use:   "jClean",
	Short: "remove json property from array",
	Long:  `print new json file without the property`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jClean called")
		jClean(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(jCleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jCleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jCleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jCleanCmd.Flags().String("key", "", "property you want remove")
	jCleanCmd.Flags().String("path", "", "path fo json file")
}
func jClean(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatal(err)
	}
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := util.ReadJsonFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var result []map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result {
		delete(v, key)
	}
	jsonOut, err := json.MarshalIndent(result, "  ", "	")
	if err != nil {
		log.Fatal(err)
	}
	ftype := util.StripFileType(path)
	f, err := os.Create(ftype + "_" + key + "_removed.json")
	if err != nil {
		f.Close()
	}
	_, err = f.Write(jsonOut)
	if err != nil {
		f.Close()
	}
}
