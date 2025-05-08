package bfc

import (
	"cast/lib"
	"os/exec"

	"github.com/spf13/cobra"
)

var MuteCmd = &cobra.Command{
	Use:   "mute",
	Short: "Mutes the default device",
	Run: func(cmd *cobra.Command, args []string) {
		cmdToRun := exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "1")

		if err := cmdToRun.Run(); err != nil {
			lib.Logger.Fatal("Failed to mute", "err", err)
		}
	},
}