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

var installCmd = &cobra.Command{
	Use:   "install",
	GroupID: "experimental",
	Short: "Installs a package",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO add support for other package managers

		if utils.IsCmdInstalled("pacman") {
			if utils.IsCmdInstalled("yay") {
				lib.Logger.Info("Installing with yay")
				cmdToRun := exec.Command("yay", "-S", args[0])

				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					lib.Logger.Fatal("Failed to install package", "err", err)
				}
			} else if utils.IsCmdInstalled("paru") {
				lib.Logger.Info("Installing with paru")
				cmdToRun := exec.Command("paru", "-S", args[0])
				
				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					lib.Logger.Fatal("Failed to install package", "err", err)
				}
			} else {
				lib.Logger.Info("Installing with pacman")
				cmdToRun := exec.Command("sudo", "pacman", "-S", args[0])
				
				cmdToRun.Stdout = os.Stdout
				cmdToRun.Stderr = os.Stderr

				if err := cmdToRun.Run(); err != nil {
					lib.Logger.Fatal("Failed to install package", "err", err)
				}
			}
		} else {
			lib.Logger.Fatal("Could not find a supported package manager")
		}
	},
}

func init() {
	if config.Config.Modules.Experimental {
		cmd.RootCmd.AddCommand(installCmd)
	}
}