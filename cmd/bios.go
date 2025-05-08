package cmd

import (
	"cast/lib"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var biosCmd = &cobra.Command{
	Use:   "bios",
	Short: "Boot into this device's BIOS/UEFI screen",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.PathExists("/sys/firmware/efi") {
			lib.Logger.Info("Rebooting to legacy BIOS from OS is not supported")
			os.Exit(0)
		} else {
			var confirm bool
			confirmation := huh.NewConfirm().
				Title("Are you sure you want to reboot to BIOS?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirm)

			if err := confirmation.Run(); err != nil {
				lib.Logger.Fatal("Failed to run the confirmation dialog", "err", err)
			}
			
			if confirm {
				cmdToRun := exec.Command("systemctl", "reboot", "--firmware-setup")

				if err := cmdToRun.Run(); err != nil {
					utils.Error("Failed to restart to BIOS.")
					os.Exit(1)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(biosCmd)
}