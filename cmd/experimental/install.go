package experimental

import (
	"cast/util"
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	GroupID: "experimental",
	Short: "Installs a package",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO add support for other package managers

		if utils.IsCmdInstalled("pacman") {
			if utils.IsCmdInstalled("yay") {
				util.Logger.Info("Installing with yay")
				cmdToRun := exec.Command("yay", "-S", args[0])

				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					util.Logger.Fatal("Failed to install package", "err", err)
				}
			} else if utils.IsCmdInstalled("paru") {
				util.Logger.Info("Installing with paru")
				cmdToRun := exec.Command("paru", "-S", args[0])
				
				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					util.Logger.Fatal("Failed to install package", "err", err)
				}
			} else {
				util.Logger.Info("Installing with pacman")
				cmdToRun := exec.Command("sudo", "pacman", "-S", args[0])
				
				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					util.Logger.Fatal("Failed to install package", "err", err)
				}
			}
		} else {
			util.Logger.Fatal("Could not find a supported package manager")
		}
	},
}