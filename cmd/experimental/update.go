package experimental

import (
	"cast/util"
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	GroupID: "experimental",
	Short: "Update system",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO add support for other package managers

		if utils.IsCmdInstalled("pacman") {
			cmdToRun := exec.Command("sudo", "pacman", "-Syyu")

			cmdToRun.Stdout = os.Stdout
			cmdToRun.Stderr = os.Stderr

			if err := cmdToRun.Run(); err != nil {
				util.Logger.Fatal("Failed to update system", "err", err)
			}
		} else {
			util.Logger.Fatal("Could not find a supported package manager")
		}
	},
}