package nova

import (
	"cast/cmd"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var (
	enableTailscaleFlag  bool
	disableTailscaleFlag bool

	tailscaleCmd = &cobra.Command{
		Use:   "tailscale",
		Short: "Enable or disable tailscale",
		Run: func(cmd *cobra.Command, args []string) {
			if enableTailscaleFlag && disableTailscaleFlag {
				utils.Warn("You can't use both --enable and --disable at the same time.")
				return
			}

			if enableTailscaleFlag {
				cmdToRun := exec.Command("systemctl", "enable", "--now", "tailscale")
				if err := cmdToRun.Run(); err != nil {
					utils.Error("Failed to enable and start Tailscale")
					return
				}

				utils.Done("Tailscale is now enabled")
				return
			}

			if disableTailscaleFlag {
				if enableTailscaleFlag {
					cmdToRun := exec.Command("systemctl", "disable", "--now", "tailscale")
					if err := cmdToRun.Run(); err != nil {
						utils.Error("Failed to disable and stop Tailscale")
						return
					}

					utils.Done("Tailscale is now disabled")
					return
				}
			}
		},
	}
)

func init() {
	if utils.IsCurrentImage("nova") {
		tailscaleCmd.Flags().BoolVar(&enableTailscaleFlag, "enable", false, "")
		tailscaleCmd.Flags().BoolVar(&disableTailscaleFlag, "disable", false, "")
		tailscaleCmd.MarkFlagsOneRequired("enable", "disable")

		cmd.RootCmd.AddCommand(tailscaleCmd)
	}
}
