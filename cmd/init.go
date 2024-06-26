/*
Copyright © 2024 avran2002@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/avran02/sky-cli/generator_modules/commands"
	"github.com/avran02/sky-cli/generator_modules/file"
	pl "github.com/avran02/sky-cli/plugins"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project config",
	Long: `This command start interactive generation question and creates project template.
You must specify plugin name as argument. It will be used to create your project.`,
	Run:     initProject,
	Example: "sky-cli init plugin",
}

// load plugin, execute all commands and generate project
func initProject(cmd *cobra.Command, args []string) {
	pluginName := mustParceArgs(args) + ".so"
	conf := pl.LoadConf(pluginName)

	err := commands.ExecAll(conf.GetOsCommands())
	if err != nil {
		fmt.Println("can't execute commands:", err)
		os.Exit(1)
	}

	fs := conf.GetVirtualFs()
	fs.Gen(".", file.AskIfNeeded, file.GetUserValues)
	fmt.Println("Project successfully generated")
}

// get plugin name from command args or exit with error code 1
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

func init() {
	rootCmd.AddCommand(initCmd)
}
