package base

import (
	"cast/util"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	GroupID: "base",
	Short: "Print the version number of cast",
	Run: func(cmd *cobra.Command, args []string) {
		util.Logger.Info("2.0")
	},
}