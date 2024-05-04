/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	plugin "github.com/avran02/sky-cli/plugins"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of available plugins",
	Run:   getPluginsList,
}

func getPluginsList(cmd *cobra.Command, args []string) {
	pluginsList := plugin.MustGetPluginNames()
	fmt.Println("Available plugins:")
	for _, pluginName := range pluginsList {
		fmt.Println(pluginName[:len(pluginName)-3])
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
