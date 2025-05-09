package base

import (
	"cast/config"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	GroupID: "base",
	Short: "Show the location of the config",
	Run: func(cmd *cobra.Command, args []string) {
		config.FindConfig()
	},
}