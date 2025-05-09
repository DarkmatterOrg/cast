package fixes

import (
	"cast/util"
	"os"
	"os/exec"
	"time"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var GmodCmd = &cobra.Command{
	Use:   "gmod",
	Short: "Patch GMod's 64-bit beta to work properly on Linux (https://github.com/solsticegamestudios/GModCEFCodecFix)",
	Run: func(cmd *cobra.Command, args []string) {
		utils.StartSpinner()

		util.Logger.Info("Creating /tmp/patch-gmod")
		if err := os.MkdirAll("/tmp/patch-gmod", os.ModePerm); err != nil {
			util.Logger.Fatal("Failed to create /tmp/patch-gmod", "err", err)
		}

		util.Logger.Info("Downloading the gmod patch")
		cmdToRun := exec.Command("sh", "-c", "wget $(curl -s https://api.github.com/repos/solsticegamestudios/GModCEFCodecFix/releases/latest | jq -r '.assets[] | select(.name | test(\"GModCEFCodecFix-Linux\")) | .browser_download_url') -P /tmp/patch-gmod")
		if err := cmdToRun.Run(); err != nil {
			util.Logger.Fatal("Failed to get the gmod patch", "err", err)
		}

		util.Logger.Info("Giving the patch run permissions")
		cmdToRun = exec.Command("chmod", "+x", "/tmp/patch-gmod/GModCEFCodecFix-Linux")
		if err := cmdToRun.Run(); err != nil {
			util.Logger.Fatal("Failed to give the patch run permissions", "err", err)
		}

		util.Logger.Info("Patching gmod")
		cmdToRun = exec.Command("/tmp/patch-gmod/GModCEFCodecFix-Linux")
		if err := cmdToRun.Run(); err != nil {
			util.Logger.Fatal("Failed to run the patch", "err", err)
		}

		time.Sleep(4 * time.Second)

		util.Logger.Info("Removing the patch")
		os.RemoveAll("/tmp/patch-gmod")

		utils.StopSpinner()
		util.Logger.Success("Fixed gmod")
	},
}