package horizon

import (
	"cast/cmd"
	"os"
	"os/exec"

	"github.com/darkmatterorg/orbit/utils"
	"github.com/spf13/cobra"
)

var (
	distroBoxUbuntuCmd = &cobra.Command{
		Use:   "dbox-ubuntu",
		Short: "Create a Ubuntu distrobox container",
		Run: func(cmd *cobra.Command, args []string) {
			if !utils.IsCmdInstalled("distrobox") {
				utils.Error("Distrobox isn't installed")
				return
			}

			cmdToRun := exec.Command("distrobox-create", "--image", "quay.io/toolbx/ubuntu-toolbox:latest", "--name", "ubuntu")
			cmdToRun.Stdout = os.Stdout
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Wasn't able to run distrobox-create")
				return
			}
		},
	}

	dboxUbuntuPkgsFlag string

	distroBoxUbuntuPkgsCmd = &cobra.Command{
		Use:   "dbox-ubuntu-with-pkgs",
		Short: "Create a Ubuntu distrobox container",
		Run: func(cmd *cobra.Command, args []string) {
			if !utils.IsCmdInstalled("distrobox") {
				utils.Error("Distrobox isn't installed")
				return
			}

			cmdToRun := exec.Command("distrobox-create", "--image", "quay.io/toolbx/ubuntu-toolbox:latest", "--name", "ubuntu", "--additional-packages", "'"+dboxUbuntuPkgsFlag+"'")
			cmdToRun.Stdout = os.Stdout
			if err := cmdToRun.Run(); err != nil {
				utils.Error("Wasn't able to run distrobox-create")
				return
			}
		},
	}
)

func init() {
	if utils.IsCurrentImage("horizon") {
		cmd.RootCmd.AddCommand(distroBoxUbuntuCmd)
		cmd.RootCmd.AddCommand(distroBoxUbuntuPkgsCmd)

		distroBoxUbuntuPkgsCmd.Flags().StringVar(&dboxUbuntuPkgsFlag, "pkgs", "", "List of packages to install, example: --pkgs=\"pkgs1 pkgs2 pkg3\"")
		distroBoxUbuntuPkgsCmd.MarkFlagRequired("pkgs")
	}
}
