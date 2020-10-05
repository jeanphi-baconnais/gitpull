package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Git-pull",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Git-pull v1.1")
	},
}