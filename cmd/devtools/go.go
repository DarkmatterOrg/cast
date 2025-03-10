package devtools

import (
	"cast/cmd"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Install or Remove go",
	Args: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().NFlag() == 0 {
			return fmt.Errorf("please use either --install or --remove flag")
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
			installGo(verboseFlag)
		} else if removeFlag {
			removeGo(verboseFlag)
		}
	},
}

func init() {
	cmd.DevToolsCmd.AddCommand(goCmd)
}

func installGo(verboseFlag bool) {
	if utils.IsCmdInstalled("go") {
		utils.Notice("Go is already installed, nothing to do.")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		utils.Error("Unable to get home directory")
		os.Exit(1)
	}

	localDir := path.Join("/usr", "local")
	tmpDir := path.Join(homeDir, "go-tmp")
	
	removeDirIfExist(tmpDir)
	os.MkdirAll(tmpDir, os.ModePerm)
	cmdToRun := exec.Command("sudo", "rm", "-rf", localDir+"/go")
	cmdToRun.Stdout = os.Stdout
	if err := cmdToRun.Run(); err != nil {
		utils.Error("Was unable to remove go")
		os.Exit(1)
	}

	utils.Info("Downloading Go...")
	cmdToRun = exec.Command("wget", "https://go.dev/dl/go1.24.1.linux-amd64.tar.gz", "-O", path.Join(tmpDir, "go.tar.gz"))
	if verboseFlag {
		cmdToRun.Stdout = os.Stdout
	} else {
		utils.StartSpinner()
	}
	if err := cmdToRun.Run(); err != nil {
		utils.StopSpinner()
		utils.Error("Failed to download Go")
		os.Exit(1)
	}
	if !verboseFlag {
		utils.StopSpinner()
	}
	utils.Done("Downloaded Go")

	utils.Info("Extracting the tar...")
	cmdToRun = exec.Command("sudo", "tar", "-xzf", path.Join(tmpDir, "go.tar.gz"), "-C", localDir)
	cmdToRun.Stdout = os.Stdout
	if err := cmdToRun.Run(); err != nil {
		fmt.Println(err)
		utils.StopSpinner()
		utils.Error("Failed to extract the tar")
		os.Exit(1)
	}
	if !verboseFlag {
		utils.StopSpinner()
	}
	utils.Done("Extracted the tar")

	profileFile := path.Join(homeDir, ".profile")

	if !utils.DoesFileContain("export PATH=$PATH:/usr/local/go/bin", profileFile) {
		utils.AppendFile("export PATH=$PATH:/usr/local/go/bin\n", profileFile)
	}

	removeDirIfExist(tmpDir)
	utils.Done("Go is installed, please relogin or source ~/.profile to see the changes.")
}

func removeGo(verboseFlag bool) {
	if !utils.IsCmdInstalled("go") {
		utils.Notice("Go is not installed, nothing todo.")
		return
	}

	localDir := path.Join("/usr", "local")
	cmdToRun := exec.Command("sudo", "rm", "-rf", localDir+"/go")
	cmdToRun.Stdout = os.Stdout
	if err := cmdToRun.Run(); err != nil {
		utils.Error("Was unable to remove go")
		os.Exit(1)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		utils.Error("Unable to get home directory")
		os.Exit(1)
	}

	profileFile := path.Join(homeDir, ".profile")

	if utils.DoesFileContain("export PATH=$PATH:/usr/local/go/bin", profileFile) {
		cmdToRun := exec.Command("sed", "-i", "/export PATH=$PATH:\\/usr\\/local\\/go\\/bin/d", profileFile)
		if err := cmdToRun.Run(); err != nil {
			utils.Error("Was unable to remove the line from the file")
			os.Exit(1)
		}
	}

	utils.Done("Go got removed, please relogin or source ~/.profile to see the changes.")
}

func removeDirIfExist(dir string) {
	if utils.PathExists(dir) {
		os.RemoveAll(dir)
	}
}