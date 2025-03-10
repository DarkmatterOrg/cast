package cmd

import (
	"github.com/spf13/cobra"
)

//! DON'T ADD ANY SUBCOMMANDS AND STUFF HERE, CREATE A FILE IN fixes/ FOR IT.
//! CHECK THE CURRENT FILES IN THERE FOR AN EXAMPLE
//!
//! IF YOU WANT TO MAKE ONE SPECIFIC FOR AN IMAGE, IT GOES IN THE IMAGE FOLDER (horizon/nova/umbra) LIKE NORMAL.
//! JUST BE SURE TO INIT WITH cmd.FixCmd.AddCommand()

var (
	FixCmd = &cobra.Command{
		Use:   "fix",
		Short: "Different fixes for various things",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	// ATM THE FIXES FOR THESE DOESN'T REALLY WORK, WILL BE IN HERE FOR THE TIME BEING

	// fixDiscord = &cobra.Command{
	// 	Use:   "discord",
	// 	Short: "Fix Discord flatpak RPC",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fixRpc("com.discordapp.Discord", "discord", "Discord")
	// 	},
	// }

	// fixVesktop = &cobra.Command{
	// 	Use:   "vesktop",
	// 	Short: "Fix Vesktop flatpak RPC",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fixRpc("dev.vencord.Vesktop", "vesktop", "Vesktop")
	// 	},
	// }
)

func init() {
	RootCmd.AddCommand(FixCmd)
	// fixCmd.AddCommand(fixDiscord)
	// fixCmd.AddCommand(fixVesktop)
}

// func fixRpc(longClient string, shortClient string, prettyname string) {
// 	configDir, err := os.UserConfigDir()
// 	if err != nil {
// 		utils.Error("Unable to get home directory")
// 		os.Exit(1)
// 	}

// 	tmpDir := path.Join(configDir, "user-tmpfiles.d")
// 	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
// 		utils.Error("Failed to create: " + tmpDir)
// 		os.Exit(1)
// 	}

// 	if shortClient == "vesktop" {
// 		if err := utils.WriteFile(path.Join(tmpDir, "discord-rpc.conf"), "L %t/discord-ipc-0 - - - - .flatpak/dev.vencord.Vesktop/xdg-run/discord-ipc-0"); err != nil {
// 			utils.Error("Was unable to create and write to file")
// 			os.Exit(1)
// 		}
// 	} else if shortClient == "discord" {
// 		if err := utils.WriteFile(path.Join(tmpDir, "discord-rpc.conf"), "L %t/discord-ipc-0 - - - - app/com.discordapp.Discord/discord-ipc-0"); err != nil {
// 			utils.Error("Was unable to create and write to file")
// 			os.Exit(1)
// 		}
// 	}

// 	cmdToRun := exec.Command("sudo", "systemctl", "enable", "--user", "--now", "systemd-tmpfiles-setup.service")
// 	if err := cmdToRun.Run(); err != nil {
// 		utils.Error("Was unable to enable systemd-tmpfiles-setup.service")
// 		os.Exit(1)
// 	}

// 	if utils.PathExists("/run/user/1000/discord-ipc-0") {
// 		os.Remove("/run/user/1000/discord-ipc-0")
// 	}

// 	if err := os.Symlink("/run/user/1000/.flatpak/"+longClient+"/xdg-run/discord-ipc-0", "/run/user/1000/discord-ipc-0"); err != nil {
// 		utils.Error("Was unable to create symlink")
// 		os.Exit(1)
// 	}

// 	utils.Notice("Will need to use sudo...")

// 	cmdToRun = exec.Command("sudo", "flatpak", "override", "--filesystem=xdg-run/discord-ipc-*")
// 	if err := cmdToRun.Run(); err != nil {
// 		utils.Error("Was unable to override discord-ipc")
// 		os.Exit(1)
// 	}

// 	cmdToRun = exec.Command("sudo", "flatpak", "override", "--filesystem=xdg-run/.flatpak/"+longClient+":create")
// 	if err := cmdToRun.Run(); err != nil {
// 		utils.Error("Was unable to override " + prettyname)
// 		os.Exit(1)
// 	}

// 	utils.Info("Adding fix to autostart...")

// 	autostartDir := path.Join(configDir, "autostart")
// 	if err := os.MkdirAll(autostartDir, os.ModePerm); err != nil {
// 		utils.Error("Unable to make: " + autostartDir)
// 		os.Exit(1)
// 	}

// 	if err := utils.WriteFile(path.Join(autostartDir, shortClient+"-rpc.desktop"), "[Desktop Entry]\nCategories=Utility;\nComment=Symlink for "+prettyname+" RPC\nIcon="+shortClient+"\nName="+prettyname+"RPC\nStartupNotify=true\nTerminal=false\nType=Application\nExec=ln -s /run/user/1000/.flatpak/"+longClient+"/xdg-run/discord-ipc-0 /run/user/1000/discord-ipc-0"); err != nil {
// 		utils.Error("Was unable to create and write to file")
// 		os.Exit(1)
// 	}

// 	utils.Done("Fixed " + prettyname + " RPC")
// }
