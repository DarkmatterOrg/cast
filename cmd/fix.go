package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var (
	fixCmd = &cobra.Command{
		Use:   "fix",
		Short: "Different fixes for various things",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	protonHangCmd = &cobra.Command{
		Use:   "proton-hang",
		Short: "Kills all processes related to wine and proton.",
		Long:  "Kills all processes related to wine and proton. This forces it to restart next time you launch the game (you might still have to press STOP in steam to kill the game binary)",
		Run: func(cmd *cobra.Command, args []string) {
			protonCore := []string{
				"pv-bwrap",
				"pressure-vessel",
				"reaper",
				"explorer.exe",
				"rpcss.exe",
				"plugplay.exe",
				"services.exe",
				"svchost.exe",
				"winedevice.exe",
				"winedevice.exe",
				"wineserver",
			}

			for _, core := range protonCore {
				if err := exec.Command("killall", "-9", core); err != nil {
					utils.Warn("Was unable to kill: " + core)
				}
			}

			utils.Done("Fixed proton hang")
		},
	}

	gmodCmd = &cobra.Command{
		Use:   "gmod",
		Short: "Patch GMod's 64-bit beta to work properly on Linux (https://github.com/solsticegamestudios/GModCEFCodecFix)",
		Run: func(cmd *cobra.Command, args []string) {
			utils.StartSpinner()
			if err := os.MkdirAll("/tmp/patch-gmod", os.ModePerm); err != nil {
				utils.Error("Failed to create /tmp/patch-gmod")
				os.Exit(1)
			}

			cmdToRun := exec.Command("sh", "-c", "wget $(curl -s https://api.github.com/repos/solsticegamestudios/GModCEFCodecFix/releases/latest | jq -r '.assets[] | select(.name | test(\"GModCEFCodecFix-Linux\")) | .browser_download_url') -P /tmp/patch-gmod")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to get the gmod patch")
				os.Exit(1)
			}

			cmdToRun = exec.Command("chmod", "+x", "/tmp/patch-gmod/GModCEFCodecFix-Linux")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to give the patch run permissions")
				os.Exit(1)
			}

			cmdToRun = exec.Command("/tmp/patch-gmod/GModCEFCodecFix-Linux")
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Failed to run the patch")
				os.Exit(1)
			}

			time.Sleep(4 * time.Second)

			os.RemoveAll("/tmp/patch-gmod")

			utils.StopSpinner()
			utils.Done("Fixed gmod")
		},
	}

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
	RootCmd.AddCommand(fixCmd)

	fixCmd.AddCommand(protonHangCmd)
	fixCmd.AddCommand(gmodCmd)
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
