package bfc

import (
	"cast/util"
	"os/exec"

	"github.com/spf13/cobra"
)

var UnmuteCmd = &cobra.Command{
	Use:   "unmute",
	Short: "Unmutes the default device",
	Run: func(cmd *cobra.Command, args []string) {
		cmdToRun := exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "0")

		if err := cmdToRun.Run(); err != nil {
			util.Logger.Fatal("Failed to unmute", "err", err)
		}
	},
}