package fixes

import (
	"cast/cmd"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var protonHangCmd = &cobra.Command{
	Use:   "proton-hang",
	Short: "Kills all processes related to wine and proton.",
	Long:  "Kills all processes related to wine and proton. This forces it to restart next time you launch the game (you might still have to press STOP in steam to kill the game binary)",
	Run: func(cmd *cobra.Command, args []string) {
		protonCore := []string{
			"pv-bwrap",
			"pressure-vessel",
			"reaper",
			"explorer.exe",
			"rpcss.exe",
			"plugplay.exe",
			"services.exe",
			"svchost.exe",
			"winedevice.exe",
			"winedevice.exe",
			"wineserver",
		}

		for _, core := range protonCore {
			if err := exec.Command("killall", "-9", core); err != nil {
				utils.Warn("Was unable to kill: " + core)
			}
		}

		utils.Done("Fixed proton hang")
	},
}

func init() {
	cmd.FixCmd.AddCommand(protonHangCmd)
}
