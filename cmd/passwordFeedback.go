package cmd

import (
	"os"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	enableFlag  bool
	disableFlag bool

	passwordFeedbackCmd = &cobra.Command{
		Use:   "password-feedback",
		Short: "Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled",
		Run: func(cmd *cobra.Command, args []string) {
			if !utils.IsRoot() {
				utils.Notice("Please run this command with sudo")
				return
			}

			if enableFlag && disableFlag {
				utils.Warn("You can't use both --enable and --disable at the same time.")
				return
			}

			if enableFlag {
				if err := utils.WriteFile("/etc/sudoers.d/enable-pwfeedback", "Defaults pwfeedback"); err != nil {
					utils.Error("Unable to write the enable-pwfeedback file")
					return
				}

				enabled := color.New(color.FgGreen, color.Bold).SprintFunc()
				utils.Notice("Password feedback is now " + enabled("enabled") + "! Restart terminal to see changes")
			}

			if disableFlag {
				os.Remove("/etc/sudoers.d/enable-pwfeedback")

				disabled := color.New(color.FgRed, color.Bold).SprintFunc()
				utils.Notice("Password feedback is now " + disabled("disabled") + "! Restart terminal to see changes")
			}
		},
	}
)

func init() {
	passwordFeedbackCmd.Flags().BoolVar(&enableFlag, "enable", false, "Enables password feedback")
	passwordFeedbackCmd.Flags().BoolVar(&disableFlag, "disable", false, "Disables password feedback")
	passwordFeedbackCmd.MarkFlagsOneRequired("enable", "disable")

	RootCmd.AddCommand(passwordFeedbackCmd)
}
