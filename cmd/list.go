/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	plugin "github.com/avran02/sky-cli/plugin_loader"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: getPluginsList,
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
