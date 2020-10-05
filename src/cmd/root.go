package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"jpbaconnais/git-pull/git"
	"jpbaconnais/git-pull/utils"
	"os"
	"jpbaconnais/git-pull/props"
	"strings"
)

var repos = make(map[int]string)

var root = &cobra.Command{
	Use:   "Git-pull",
	Short: "Git-pull",
	Long: `Tools which allows to refresh all your repo Git`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("🙌 Git pull ... ")

		// verif properties file
		path := props.VerifPropertiesFile()

		fmt.Println("A properties file is found, let's start git pull ")

		getRepositoriesGit(path)

		for _, repo := range repos {
			git.LaunchGitOperations(repo)
		}
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


func getRepositoriesGit(root string) {

	directories, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println("⁉️  🍄 Oh, this directory isn't available : ", root, " - err ", err)
	}

	//isGit := false
	for _, dir := range directories {
		//fmt.Println("   dir ", dir.Name())
		if dir.IsDir() {
			if strings.Compare(dir.Name(), ".git") == 0 {
				//fmt.Println("   -> repo git")
				//isGit = true
				repos[len(repos)+1] = strings.ReplaceAll(root, "//", "/")
				break
			}
		}
	}

	for _, dir := range directories {
		if dir.IsDir() && utils.IsRealDir(dir.Name()) {
			getRepositoriesGit(root + "/" + dir.Name())
		}
	}

}