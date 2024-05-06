/*
Copyright © 2024 avran2002@gmail.com
*/
package cmd

import (
	"fmt"

	plugin "github.com/avran02/sky-cli/plugins"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of available plugins",
	Long: `This command get list of available plugins. 
It search in ~/.config/sky-cli/plugins directory and cut 3 last symbols from filename.`,
	Run: getPluginsList,
}

// get list of available plugins from ~/.config/sky-cli/plugins
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
