package fixes

import (
	"cast/cmd"
	"os"
	"os/exec"
	"time"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var gmodCmd = &cobra.Command{
		Use:   "gmod",
		Short: "Patch GMod's 64-bit beta to work properly on Linux (https://github.com/solsticegamestudios/GModCEFCodecFix)",
		Run: func(cmd *cobra.Command, args []string) {
			utils.StartSpinner()
			if err := os.MkdirAll("/tmp/patch-gmod", os.ModePerm); err != nil {
				utils.Error("Failed to create /tmp/patch-gmod")
				os.Exit(1)
			}

			cmdToRun := exec.Command("sh", "-c", "wget $(curl -s https://api.github.com/repos/solsticegamestudios/GModCEFCodecFix/releases/latest | jq -r '.assets[] | select(.name | test(\"GModCEFCodecFix-Linux\")) | .browser_download_url') -P /tmp/patch-gmod")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to get the gmod patch")
				os.Exit(1)
			}

			cmdToRun = exec.Command("chmod", "+x", "/tmp/patch-gmod/GModCEFCodecFix-Linux")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to give the patch run permissions")
				os.Exit(1)
			}

			cmdToRun = exec.Command("/tmp/patch-gmod/GModCEFCodecFix-Linux")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to run the patch")
				os.Exit(1)
			}

			time.Sleep(4 * time.Second)

			os.RemoveAll("/tmp/patch-gmod")

			utils.StopSpinner()
			utils.Done("Fixed gmod")
		},
	}

func init() {
	cmd.FixCmd.AddCommand(gmodCmd)
}
