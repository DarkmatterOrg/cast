package umbra

import (
	"cast/cmd/umbra/bfc"

	"github.com/spf13/cobra"
)

var BitfocusCompanionCmd = &cobra.Command{
	Use:   "bfc",
	GroupID: "umbra",
	TraverseChildren: true,
	Short: "Commands to use from a stream deck with Bitfocus Companion",
}



func init() {
	BitfocusCompanionCmd.AddCommand(bfc.VolumeCmd)
	BitfocusCompanionCmd.AddCommand(bfc.MuteCmd)
	BitfocusCompanionCmd.AddCommand(bfc.UnmuteCmd)
}