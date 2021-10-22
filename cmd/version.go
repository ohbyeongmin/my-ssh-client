package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "v1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of My-SSH-Client",
	Long:  `All software has versions. This is MSC's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("My-SSH-Client Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
