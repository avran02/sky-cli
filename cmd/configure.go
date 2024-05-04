/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

func configureSkyCli(cmd *cobra.Command, _ []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mustCreateDir(home + "/.config")
	mustCreateDir(home + "/.config/sky-cli")
	mustCreateDir(home + "/.config/sky-cli/plugins")
	plugins.Install("github.com/avran02/plugin@v1.0.1")
	// plugins.Install("github.com/avran02/sky-cli-plugins@v0.0.1/go-grpc")
}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
