package cmd

import (
	"cast/config"
	"cast/lib"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
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
	RootCmd.SetErrPrefix(errorTitle("ERROR"))


	if config.Config.Modules.Experimental {
		RootCmd.AddGroup(&cobra.Group{ID: "experimental", Title: "Experimental"})
	}
}

func Execute() {
	if config.Config.Modules.Experimental {
		if config.Config.Insult {
			lib.ImportantWarn("Either you're fucking stupid or just curious. Either way; experimental features are enabled, use at your own risk")
		} else {
			lib.ImportantWarn("Experimental features are enabled, use at your own risk")
		}
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}