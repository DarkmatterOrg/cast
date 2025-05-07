package experimental

import (
	"cast/cmd"
	"cast/config"
	"cast/lib"
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "*EXPERIMENTAL* Update system",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO add support for other package managers

		if utils.IsCmdInstalled("pacman") {
			cmdToRun := exec.Command("sudo", "pacman", "-Syyu")

			cmdToRun.Stdout = os.Stdout
			cmdToRun.Stderr = os.Stderr

			if err := cmdToRun.Run(); err != nil {
				lib.Logger.Fatal("Failed to update system", "err", err)
			}
		} else {
			lib.Logger.Fatal("Could not find a supported package manager")
		}
	},
}

func init() {
	if config.Config.Modules.Experimental {
		cmd.RootCmd.AddCommand(updateCmd)
	}
}