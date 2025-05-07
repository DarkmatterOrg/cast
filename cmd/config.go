package cmd

import (
	"cast/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show the location of the config",
	Run: func(cmd *cobra.Command, args []string) {
		config.FindConfig()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
}