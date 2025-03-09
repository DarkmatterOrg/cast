package cmd

import (
	"os/exec"
	"strings"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	statusFlag            bool
	enableAutoUpdateFlag  bool
	disableAutoUpdateFlag bool

	autoUpdateCmd = &cobra.Command{
		Use:   "auto-update",
		Short: "Toggle the auto-updater",
		Run: func(cmd *cobra.Command, args []string) {
			timeFileName := getTimeFileName()
			if timeFileName == "" {
				utils.Error("No corect image type")
				return
			}

			if statusFlag {
				cmdToRun := exec.Command("systemctl", "is-enabled", timeFileName)

				output, _ := cmdToRun.Output()

				outputStatus := color.New(color.FgRed, color.Bold).SprintFunc()
				outputStatusText := outputStatus("disabled")

				if strings.Contains(string(output), "enabled") {
					outputStatus = color.New(color.FgGreen, color.Bold).SprintFunc()
					outputStatusText = outputStatus("enabled")
				}

				utils.Info("Automatic updates are currently: " + outputStatusText)
			}

			if enableAutoUpdateFlag && disableAutoUpdateFlag {
				utils.Warn("You can't use both --enable and --disable at the same time.")
				return
			}

			if enableAutoUpdateFlag {
				cmdToRun := exec.Command("systemctl", "enable", timeFileName)
				if err := cmdToRun.Run(); err != nil {
					utils.Error("Failed to enable the auto update timer")
					return
				}

				enabled := color.New(color.FgGreen, color.Bold).SprintFunc()
				utils.Notice("Auto Updater have been " + enabled("enabled"))
			}

			if disableAutoUpdateFlag {
				cmdToRun := exec.Command("systemctl", "disable", timeFileName)
				if err := cmdToRun.Run(); err != nil {
					utils.Error("Failed to disable the auto update timer")
					return
				}

				disabled := color.New(color.FgRed, color.Bold).SprintFunc()
				utils.Notice("Auto Updater have been " + disabled("disabled"))
			}
		},
	}
)

func init() {
	autoUpdateCmd.Flags().BoolVarP(&statusFlag, "status", "s", false, "Get the current status of the updater")
	autoUpdateCmd.Flags().BoolVarP(&enableAutoUpdateFlag, "enable", "e", false, "")
	autoUpdateCmd.Flags().BoolVarP(&disableAutoUpdateFlag, "disable", "d", false, "")
	autoUpdateCmd.MarkFlagsOneRequired("status", "enable", "disable")

	RootCmd.AddCommand(autoUpdateCmd)
}

func getTimeFileName() string {
	imageType := utils.Getimagetype()

	if strings.Contains(imageType, "umbra") {
		return "umbra-update.timer"
	} else if strings.Contains(imageType, "nova") {
		return "nova-update.timer"
	} else if strings.Contains(imageType, "aster") || strings.Contains(imageType, "arcturus") {
		return "horison-update.timer"
	}
	return ""
}
