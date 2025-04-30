package cmd

import (
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var cleanSystemCmd = &cobra.Command{
	Use:   "clean-system",
	Short: "Clean up old up unused podman images, volumes, flatpak packages and rpm-ostree content",
	Run: func(cmd *cobra.Command, args []string) {

		utils.Info("Cleaning old podman images")
		if err := exec.Command("podman", "image", "prune", "-af").Run(); err != nil {
			utils.Error("Failed to prune podman images")
		} else {
			utils.Done("Pruned podman images")
		}

		utils.Info("Cleaning old podman volumes")
		if err := exec.Command("podman", "volume", "prune", "-f").Run(); err != nil {
			utils.Error("Failed to prune podman volumes")
		} else {
			utils.Done("Pruned podman volumes")
		}

		utils.Info("Cleaning unused flatpak packages")
		if err := exec.Command("flatpak", "uninstall", "--unused").Run(); err != nil {
			utils.Error("Failed to uninstall unused flatpaks")
		} else {
			utils.Done("Uninstalled unused flatpaks")
		}

		utils.Info("Cleaning up rpm-ostree")
		if err := exec.Command("rpm-ostree", "cleanup", "-bm").Run(); err != nil {
			utils.Error("Failed to cleanup rpm-ostree")
		} else {
			utils.Done("Cleaned up rpm-ostree")
		}
	},
}

func init() {
	RootCmd.AddCommand(cleanSystemCmd)
}
