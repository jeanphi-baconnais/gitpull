package cmd

import (
	"github.com/spf13/cobra"
	"jpbaconnais/git-pull/props"
)

func init() {
	root.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Git-pull",
	Long:  `git-pull init ...`,
	Run: func(cmd *cobra.Command, args []string) {
		props.InitTool()
	},
}
