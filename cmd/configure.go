/*
Copyright © 2024 Avranenko Andrey avran2002@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/avran02/sky-cli/plugins"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure project config",
	Run:   configureSkyCli,
}

// Create ~/.config/sky-cli/plugins directory and download plugin generator
func configureSkyCli(cmd *cobra.Command, args []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mustCreateDir(home + "/.config")
	mustCreateDir(home + "/.config/sky-cli")
	mustCreateDir(home + "/.config/sky-cli/plugins")
	plugins.Install("github.com/avran02/plugin@v1.0.1")
}

// if dir exists - do nothing, if not - create or exit with error
func mustCreateDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, 0755) //nolint
		if err != nil {
			fmt.Println("can't create dir")
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
