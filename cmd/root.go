package cmd

import (
	"cast/cmd/base"
	"cast/cmd/experimental"
	"cast/cmd/umbra"
	"cast/config"
	"cast/lib"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cast",
		Short: "Cast is a swiss army knife.",
		Long:  "Go Cast Go is a swiss army knife, usable on any Distro.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 && cmd.Flags().NFlag() == 0 {
				_ = cmd.Help()
				return
			}
		},
	}
)

func init() {
	config.LoadConfig()
	// RootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	errorTitle := color.New(color.FgRed, color.Bold).SprintFunc()
	rootCmd.SetErrPrefix(errorTitle("ERROR"))

	rootCmd.AddGroup(&cobra.Group{ID: "base", Title: "Base Commands"})
	rootCmd.AddCommand(base.VersionCmd)
	rootCmd.AddCommand(base.PasswordFeedbackCmd)
	rootCmd.AddCommand(base.ConfigCmd)
	rootCmd.AddCommand(base.BiosCmd)


	if config.Config.Modules.Experimental {
		rootCmd.AddGroup(&cobra.Group{ID: "experimental", Title: "Experimental Commands"})
		rootCmd.AddCommand(experimental.InstallCmd)
		rootCmd.AddCommand(experimental.UpdateCmd)
	}

	if config.Config.Modules.Umbra {
		rootCmd.AddGroup(&cobra.Group{ID: "umbra", Title: "Umbra Commands"})
		rootCmd.AddCommand(umbra.BitfocusCompanionCmd)
	}

	if config.Config.Modules.Fixes {
		rootCmd.AddCommand(FixCmd)
	}
}

func Execute() {
	if config.Config.Modules.Experimental {
		if config.Config.Insult {
			lib.Logger.ImportantWarn("Either you're fucking stupid or just curious. Either way; experimental features are enabled, use at your own risk")
		} else {
			lib.Logger.ImportantWarn("Experimental features are enabled, use at your own risk")
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}