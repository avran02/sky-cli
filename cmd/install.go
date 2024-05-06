/*
Copyright © 2024 avran2002@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/avran02/sky-cli/plugins"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugin from url",
	Long: `This command install plugin from url. You must specify url as argument. 
WARNING! This command will override plugin if it exists.
WARNING! You must specify plugin version and you can't use @latest tg at the moment.`,
	Run:     installPlugin,
	Example: "sky-cli install github.com/avran02/plugin@v1.0.1",
}

// download plugin from url and compile it as .so file in ~/.config/sky-cli/plugins
func installPlugin(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		fmt.Println("wrong number of arguments. Excepted only plugin url, but got", len(args), "args")
	}
	plugins.Install(args[0])
	fmt.Println("plugin", args[0], "was installed")
}

func init() {
	rootCmd.AddCommand(installCmd)
}
