package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	RootCmd.SetHelpFunc(customizeHelp)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func customizeHelp(cmd *cobra.Command, args []string) {
	// Modify the "Usage" section
	usageLine := cmd.UseLine()
	parts := strings.SplitN(usageLine, " ", 2)
	if len(parts) > 1 {
		usageLine = parts[0] + " " + bold(parts[1]) // Bold only arguments
	} else {
		usageLine = bold(usageLine)
	}

	// Print the custom help message
	fmt.Println(bold("Go Cast Go is a user utility designed to be used for Darkmatter's different images."))
	fmt.Println("\n" + bold("Usage") + ":")
	fmt.Println("  " + bold(usageLine))
	fmt.Println("\n" + bold("Available Commands"+":"))

	// Print aligned commands with fixed column width
	for _, c := range cmd.Commands() {
		if !c.Hidden {
			fmt.Printf("  %-*s %s\n", columnWidth, bold(c.Name()), c.Short)
		}
	}

	fmt.Println("\nFlags:")

	// Print aligned flags with fixed column width
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		flagText := "--" + f.Name
		if f.Shorthand != "" {
			flagText = "-" + f.Shorthand + ", " + flagText
		}
		fmt.Printf("  %-*s %s\n", columnWidth, bold(flagText), f.Usage)
	})

	fmt.Println("\nUse " + bold("cast [command] --help") + " for more information about a command.")
}

const columnWidth = 30

func formatWithPadding(s string) string {
	return fmt.Sprintf("%-*s", columnWidth, s)
}

func bold(input string) string {
	return "\033[1m" + input + "\033[0m" // Apply ANSI bold
}
