package cmd

import (
	"fmt"
	"os"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	versionFlag bool

	RootCmd = &cobra.Command{
		Use:   "cast",
		Short: "Cast is a user utility for Darkmatter's images.",
		Long:  "Go Cast Go is a user utility designed to be used for Darkmatter's different images.",
		Run: func(cmd *cobra.Command, args []string) {
			if !utils.IsCurrentImage("horizon") && !utils.IsCurrentImage("nova") && !utils.IsCurrentImage("umbra") {
				utils.Error("Cast can only be ran on Darkmatter's different images.")
				return
			}

			if len(args) == 0 && cmd.Flags().NFlag() == 0 {
				_ = cmd.Help()
				return
			}

			if versionFlag {
				fmt.Println(color.MagentaString("Cast") + ": 1.0")
			}
		},
	}
)

func init() {
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "V", false, "Show version")
	RootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
