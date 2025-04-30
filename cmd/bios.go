package cmd

import (
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var biosCmd = &cobra.Command{
	Use:   "bios",
	Short: "Boot into this device's BIOS/UEFI screen",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.PathExists("/sys/firmware/efi") {
			utils.Notice("Rebooting to legacy BIOS from OS is not supported")
			return
		} else {
			//TODO Ask the user if they are sure
			cmdToRun := exec.Command("systemctl", "reboot", "--firmware-setup")

			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to restart to BIOS.")
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(biosCmd)
}
