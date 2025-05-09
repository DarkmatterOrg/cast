package bfc

import (
	"cast/util"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	increaseFlag bool
	decreaseFlag bool
	amountFlag int

	VolumeCmd = &cobra.Command{
		Use:   "volume",
		Short: "Increase or decrease volume on default device",
		Example: "cast bfc volume --increase --amount 10",
		Run: func(cmd *cobra.Command, args []string) {
			if increaseFlag && decreaseFlag {
				if util.Config.Insult {
					util.Logger.Warn("You can't use both --increase and --decrease at the same time you fucking moron!")
					os.Exit(0)
				} else {
					util.Logger.Warn("You can't use both --increase and --decrease at the same time.")
					os.Exit(0)
				}
			}

			if increaseFlag {
				cmdToRun := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", fmt.Sprintf("+%d%%", amountFlag))

				if err := cmdToRun.Run(); err != nil {
					util.Logger.Fatal("Failed to increase volume", "err", err)
				}
			}

			if decreaseFlag {
				cmdToRun := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", fmt.Sprintf("-%d%%", amountFlag))

				if err := cmdToRun.Run(); err != nil {
					util.Logger.Fatal("Failed to decrease volume", "err", err)
				}
			}
		},
	}
)

func init() {
	VolumeCmd.Flags().BoolVarP(&increaseFlag, "increase", "i", false, "Increases volume")
	VolumeCmd.Flags().BoolVarP(&decreaseFlag, "decrease", "d", false, "Decreases volume")
	VolumeCmd.Flags().IntVarP(&amountFlag, "amount", "a", 5, "Amount to increase or decrease volume by")
	VolumeCmd.MarkFlagsOneRequired("increase", "decrease")
}