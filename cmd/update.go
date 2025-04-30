package cmd

import (
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var (
	systemFlag bool
	userFlag   bool

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Run an update on user, system or both",
		Run: func(cmd *cobra.Command, args []string) {
			if !utils.IsCmdInstalled("nebula") {
				utils.Error("Nebula needs to be installed")
				return
			}

			if systemFlag {
				utils.Info("Updating system")
				updateSystem()
			}

			if userFlag {
				utils.Info("Updating user")
				updateUser()
			}
		},
	}
)

func init() {
	updateCmd.Flags().BoolVarP(&systemFlag, "system", "s", false, "")
	updateCmd.Flags().BoolVarP(&userFlag, "user", "u", false, "")
	updateCmd.MarkFlagsOneRequired("system", "user")

	RootCmd.AddCommand(updateCmd)
}

func updateSystem() {
	cmdToRun := exec.Command("sudo", "nebula", "update-system")

	cmdToRun.Stdout = os.Stdout

	if err := cmdToRun.Run(); err != nil {
		utils.Error("Failed to update the system")
		os.Exit(1)
	}
}

func updateUser() {
	cmdToRun := exec.Command("nebula", "update-system")

	cmdToRun.Stdout = os.Stdout

	if err := cmdToRun.Run(); err != nil {
		utils.Error("Failed to update the user")
		os.Exit(1)
	}
}
