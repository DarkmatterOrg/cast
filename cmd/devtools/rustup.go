package devtools

import (
	"cast/cmd"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var rustupCmd = &cobra.Command{
	Use:   "rustup",
	Short: "Install or Remove rustup",
	Run: func(cmd *cobra.Command, args []string) {
		installFlag, _ := cmd.Flags().GetBool("install")
		removeFlag, _ := cmd.Flags().GetBool("remove")
		verboseFlag, _ := cmd.Flags().GetBool("verbose")

		if installFlag && removeFlag {
			utils.Warn("You can't use both --install and --remove at the same time.")
			return
		}

		if installFlag {
			installRustup(verboseFlag)
		} else if removeFlag {
			removeRustup(verboseFlag)
		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	cmd.DevToolsCmd.AddCommand(rustupCmd)
}

func installRustup(verboseFlag bool) {
	if utils.IsCmdInstalled("rustup") {
		utils.Notice("Rustup is already installed, nothing to do.")
		return
	}

	cmdToRun := exec.Command("sh", "-c", "curl --tlsv1.2 -sSf 'https://sh.rustup.rs' | sh -s -- -y")
	utils.Info("Installing rustup...")

	if verboseFlag {
		cmdToRun.Stdout = os.Stdout
	} else {
		utils.StartSpinner()
	}

	err := cmdToRun.Run()

	if !verboseFlag {
		utils.StopSpinner()
	}

	if err != nil {
		utils.Error("Failed to install rustup")
		return
	}

	utils.Done("Rustup is installed, please resart the terminal to start using it.")
}

func removeRustup(verboseFlag bool) {
	if !utils.IsCmdInstalled("rustup") {
		utils.Notice("Rustup is not installed, nothing todo.")
		return
	}

	s := spinner.New(spinner.CharSets[43], 150*time.Millisecond)
	cmdToRun := exec.Command("rustup", "self", "uninstall", "-y")
	utils.Info("Removing rustup...")

	if verboseFlag {
		cmdToRun.Stdout = os.Stdout
	} else {
		s.Start()
	}

	err := cmdToRun.Run()

	if !verboseFlag {
		s.Stop()
	}

	if err != nil {
		utils.Error("Failed to remove rustup")
		return
	}

	utils.Done("Rustup got removed, please restart the terminal to see the changes.")
}
