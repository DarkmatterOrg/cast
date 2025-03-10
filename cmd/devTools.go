package cmd

import (
	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

//! DON'T ADD ANY SUBCOMMANDS AND STUFF HERE, CREATE A FILE IN devtools/ FOR IT.
//! CHECK THE CURRENT FILES IN THERE FOR AN EXAMPLE
//!
//! IF YOU WANT TO MAKE ONE SPECIFIC FOR AN IMAGE, IT GOES IN THE IMAGE FOLDER (horizon/nova/umbra) LIKE NORMAL.
//! JUST BE SURE TO INIT WITH cmd.DevToolsCmd.AddCommand()

var (
	installFlag bool
	removeFlag  bool
	verboseFlag bool

	DevToolsCmd = &cobra.Command{
		Use:   "dev-tool",
		Short: "Install different programming tools",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				_ = cmd.Help()
				return
			}

			if installFlag && removeFlag {
				utils.Warn("You can't use both --install and --remove at the same time.")
				return
			}
		},
	}
)

func init() {
	DevToolsCmd.PersistentFlags().BoolVarP(&installFlag, "install", "i", false, "Install the tool")
	DevToolsCmd.PersistentFlags().BoolVarP(&removeFlag, "remove", "r", false, "Remove the tool")
	DevToolsCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Get more of the outputs")

	RootCmd.AddCommand(DevToolsCmd)
}
