package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ohbyeongmin/my-ssh-client/config"
	"github.com/spf13/cobra"
)

var index int32
var serverListFile2 string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to server",
	Run: func(cmd *cobra.Command, args []string) {
		a := config.GetServer(serverListFile2, 1)
		acmd := exec.Command("ssh", fmt.Sprintf("centos@%s", a.Ip), "-i", a.Key)
		acmd.Stdin = os.Stdin
		acmd.Stdout = os.Stdout
		acmd.Stderr = os.Stderr
		acmd.Run()
	},
}

func init() {
	connectCmd.Flags().Int32VarP(&index, "index", "i", 1, "Input Server Index")
	connectCmd.Flags().StringVarP(&serverListFile2, "file", "f", "./config/server_list.yaml", "Input your server list file")
	rootCmd.AddCommand(connectCmd)
}
