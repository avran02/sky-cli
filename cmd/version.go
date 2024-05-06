/*
Copyright © 2024 avran2002@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/avran02/sky-cli/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show sky-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sky-cli version:", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
