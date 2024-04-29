/*
Copyright © 2024 avran2002@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	plugin "github.com/avran02/sky-cli/plugin_loader"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project config",
	Long: `This command initializes a new project generation config file. This file will be deleted after project generation
To execute this command, you must be in a project directory and you need to have cobra-cli installed to run following command
It takes one optional argument: your email to put in in cobra-cli generated copyright. If not provided, "unknown" will be used`,
	Run:     initProject,
	Example: "sky-cli init mail@example.com",
}

func initProject(cmd *cobra.Command, args []string) {
	pluginName := mustParceArgs(args) + ".so"
	conf := plugin.LoadConf(pluginName)
	configure := make(map[string]bool, 0)

	requiredOptions := conf.GetRequiredParams()
	for _, option := range requiredOptions {
		configure[option] = true
	}

	allOptions := conf.GetAvailalableOptions()
	for _, option := range allOptions {
		configure[option] = askIfNeeded(option)
	}

	fs := conf.GetVirtualFs(configure)
	fs.Gen(".", conf)
}

func mustParceArgs(args []string) string {
	numArgs := len(args)
	var plugin string
	switch numArgs {
	case 0:
		fmt.Println("No arguments. Expected plugin name")
		os.Exit(1)
	case 1:
		plugin = args[0]
	default:
		fmt.Println("Too many arguments. Expected only plugin name")
		os.Exit(1)
	}
	return plugin
}

func askIfNeeded(option string) bool {
	fmt.Println("is", option, "needed? [Y/n]")
	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "y", "Y", "":
		return true
	case "n", "N":
		return false
	default:
		fmt.Println("Wrong answer")
		os.Exit(1)
	}
	return false
}

func init() {
	rootCmd.AddCommand(initCmd)
}