package cmd

import (
	"fmt"

	"github.com/ohbyeongmin/my-ssh-client/config"
	"github.com/spf13/cobra"
)

var serverListFile string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print Server list",
	Run: func(cmd *cobra.Command, args []string) {
		list := config.GetServerList(serverListFile)
		for i, v := range list {
			fmt.Printf("%d : %s\n", i+1, v)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&serverListFile, "file", "f", "./config/server_list.yaml", "Input your server list file")
	rootCmd.AddCommand(listCmd)
}
