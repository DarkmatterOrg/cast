package umbra

import (
	"cast/cmd"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var (
	mcreatorVersionFlag string

	mcreatorCmd = &cobra.Command{
		Use:   "mcreator",
		Short: "Install or Remove MCreator",
		Args: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().NFlag() == 0 {
				return fmt.Errorf("Please use either --install or --remove flag")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			installFlag, _ := cmd.Flags().GetBool("install")
			removeFlag, _ := cmd.Flags().GetBool("remove")
			verboseFlag, _ := cmd.Flags().GetBool("verbose")

			if installFlag && removeFlag {
				utils.Warn("You can't use both --install and --remove at the same time.")
				return
			}

			if installFlag {
				installMcreator(verboseFlag)
			} else if removeFlag {
				removeMcreator()
			}
		},
	}
)

func init() {
	if utils.IsCurrentImage("umbra") {
		mcreatorCmd.Flags().StringVarP(&mcreatorVersionFlag, "version", "V", "", "(Optional) Possible values [2024.1, 2024.2, 2024.3, 2024.4]")

		cmd.DevToolsCmd.AddCommand(mcreatorCmd)
	}
}

func installMcreator(verboseFlag bool) {
	var downloadUrl = "https://github.com/MCreator/MCreator/releases/download/"

	switch mcreatorVersionFlag {
	case "2024.1":
		downloadUrl += "2024.1.18518/MCreator.2024.1.Linux.64bit.tar.gz"
	case "2024.2":
		downloadUrl += "2024.2.32117/MCreator.2024.2.Linux.64bit.tar.gz"
	case "2024.3":
		downloadUrl += "2024.3.42716/MCreator.2024.3.Linux.64bit.tar.gz"
	case "2024.4":
		downloadUrl += "2024.4.52410/MCreator.2024.4.Linux.64bit.tar.gz"
	default:
		downloadUrl += "2025.1.10519/MCreator.2025.1.Linux.64bit.tar.gz"
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		utils.Error("Unable to get home directory")
		os.Exit(1)
	}

	appDir := path.Join(homeDir, ".local", "share", "applications")
	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		utils.Error("Failed to create: " + appDir)
		os.Exit(1)
	}

	tarFile := path.Join(appDir, "mcreator.tar.gz")
	mcreatorDir := path.Join(appDir, "mcreator")
	mcreatorDesktopFile := path.Join(appDir, "mcreator.desktop")

	removeTarIfExist(tarFile)

	utils.Info("Downloading MCreator...")
	cmdToRun := exec.Command("wget", downloadUrl, "-O", tarFile)
	if verboseFlag {
		cmdToRun.Stdout = os.Stdout
	} else {
		utils.StartSpinner()
	}
	if err := cmdToRun.Run(); err != nil {
		utils.Error("Failed to download MCreator")
		os.Exit(1)
	}
	if !verboseFlag {
		utils.StopSpinner()
	}
	utils.Done("Downloaded MCreator")

	removeMcreatorDirIfExist(mcreatorDir)

	utils.Info("Extracting the tar...")
	os.MkdirAll(mcreatorDir, os.ModePerm)
	cmdToRun = exec.Command("tar", "-xf", tarFile, "-C", mcreatorDir, "--strip-components=1")
	if verboseFlag {
		cmdToRun.Stdout = os.Stdout
	} else {
		utils.StartSpinner()
	}
	if err := cmdToRun.Run(); err != nil {
		utils.Error("Failed to extract the tar")
		os.Exit(1)
	}
	if !verboseFlag {
		utils.StopSpinner()
	}
	utils.Done("Extracted the tar")

	removeDesktopFileIfExist(mcreatorDesktopFile)

	if err := utils.WriteFile(path.Join(appDir, "mcreator.desktop"), "[Desktop Entry]\nExec=/bin/bash -c 'cd \""+mcreatorDir+"\" && ./mcreator.sh'\nType=Application\nTerminal=false\nName=MCreator\nIcon="+mcreatorDir+"/icon.png"); err != nil {
		utils.Error("Was unable to create and write mcreator.desktop")
		os.Exit(1)
	}

	removeTarIfExist(tarFile)
	utils.Done("MCreator has been installed")
}

func removeMcreator() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		utils.Error("Unable to get home directory")
		os.Exit(1)
	}
	appDir := path.Join(homeDir, ".local", "share", "applications")
	tarFile := path.Join(appDir, "mcreator.tar.gz")
	mcreatorDir := path.Join(appDir, "mcreator")
	mcreatorDesktopFile := path.Join(appDir, "mcreator.desktop")

	removeTarIfExist(tarFile)
	removeMcreatorDirIfExist(mcreatorDir)
	removeDesktopFileIfExist(mcreatorDesktopFile)
	utils.Done("MCreator has been removed")
}

func removeTarIfExist(tarFile string) {
	if utils.PathExists(tarFile) {
		os.RemoveAll(tarFile)
	}
}

func removeMcreatorDirIfExist(mcreatorDir string) {
	if utils.PathExists(mcreatorDir) {
		os.RemoveAll(mcreatorDir)
	}
}

func removeDesktopFileIfExist(mcreatorDesktopFile string) {
	if utils.PathExists(mcreatorDesktopFile) {
		os.RemoveAll(mcreatorDesktopFile)
	}
}
