package cmd

import (
	"cast/lib"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cast",
	Run: func(cmd *cobra.Command, args []string) {
		lib.Logger.Info("2.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}