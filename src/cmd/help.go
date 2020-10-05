package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(helpCmd)
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Git-pull Help",
	Long:  `git-pull help ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Oh, you need some help. Let's me your introduce the Git pull tool 😃")
		fmt.Println("Git pull allows to make a simple git pull on all yours repositories save in one directory.")
		fmt.Println("To configure your repository directory, you can take gitpull help")
		fmt.Println("And next, you can only execute gitpull to pull all yours git repositories")
		fmt.Println("Enjoy ! 😎")
	},
}