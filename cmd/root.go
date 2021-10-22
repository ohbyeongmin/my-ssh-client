package cmd

import (
	"github.com/spf13/cobra"
)

var MySSHFile string

var rootCmd = &cobra.Command{
	Use:   "myssh",
	Short: "Welcome to My-SSH-Client!",
}

func Execute() error {
	return rootCmd.Execute()
}
