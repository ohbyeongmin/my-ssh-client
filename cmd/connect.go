package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ohbyeongmin/my-ssh-client/config"
	"github.com/spf13/cobra"
)

var index int
var serverName string

func preRun(path string) {
	preRunCmd := exec.Command("/bin/sh", path)
	fmt.Println("PreRun script Run!")
	preRunOutput, err := preRunCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(preRunOutput))
}

func connectToSSH(info config.Server) {
	server := fmt.Sprintf("%s@%s", info.Account, info.Ip)
	sshCmd := exec.Command("ssh", server, "-p", fmt.Sprintf("%d", info.Port), "-i", info.Key)
	sshCmd.Stdin = os.Stdin
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	err := sshCmd.Run()
	if err != nil {
		panic(err)
	}
}

func getServerInfo() (config.Server, error) {
	if index != -1 {
		s, err := config.GetServerByIndex(index)
		if err != nil {
			return config.Server{}, err
		}
		return s, nil
	} else if serverName != "" {
		s, err := config.GetServerByName(serverName)
		if err != nil {
			return config.Server{}, err
		}
		return s, nil
	}
	return config.Server{}, errors.New("can't not find server")
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to server",
	Run: func(cmd *cobra.Command, args []string) {
		if MySSHFile == "" {
			fmt.Println("Please input your server info file")
			return
		}
		config.SetFileConfig(MySSHFile)
		if preRunPath := config.GetPreRun(); preRunPath != "" {
			preRun(preRunPath)
		}
		server, err := getServerInfo()
		if err != nil {
			panic(err)
		}
		connectToSSH(server)
	},
}

func init() {
	connectCmd.Flags().StringVarP(&serverName, "name", "n", "", "Input Server Name")
	connectCmd.Flags().IntVarP(&index, "index", "i", -1, "Input Server Index")
	connectCmd.Flags().StringVarP(&MySSHFile, "file", "f", "", "Input your server list file")
	rootCmd.AddCommand(connectCmd)
}
