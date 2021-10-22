package cmd

import (
	"fmt"

	"github.com/ohbyeongmin/my-ssh-client/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print Server list",
	Run: func(cmd *cobra.Command, args []string) {
		if MySSHFile == "" {
			fmt.Println("Please input your server info file")
			return
		}
		config.SetFileConfig(MySSHFile)
		list := config.GetServerList()
		for i, v := range list {
			fmt.Printf("%d : %s\n", i+1, v)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&MySSHFile, "file", "f", "", "Input your server list file")
	rootCmd.AddCommand(listCmd)
}
