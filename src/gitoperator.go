package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func LaunchGitOperations(repo string) {
	fmt.Println("\n   🐢 Launch Git operations on the repo ", repo)
	branch, _ := getCurrentBranch(repo)
	gitPull(repo, branch)
}

func getCurrentBranch(repo string) (string, error) {
	fmt.Println("🤞 🧞‍♂️ Let's go to find the current branch of the repository '", repo, "'")

	// get current branch
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = repo
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("⁉️  🍄 Git status in error 🐛 on this repository ", repo, " - ", err.Error())
		return "", err
	}

	branch := string(stdout)
	branch = strings.TrimLeft(strings.Replace(branch, "\n", "", 2), " ")

	fmt.Println("👉 🧞‍♂️ Git status ", repo, " : ", branch)

	return branch, nil

}

func gitPull(repo string, branch string) {
	fmt.Println("🤞 🧚‍♀️ Let's go to Git pull the repository '", repo, "' on the branch '", branch, "'")

	// get current branch
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = repo
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("⁉️  🍄 Git pull in error 🐛 on this repository ", repo, " - ", err.Error())
		return
	}

	fmt.Print("👉 🧚‍♀️ Git pull ", repo, " : ", string(stdout))

}
