package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


var commandHelp = "help"
var commandInit = "init"
var repos = make(map[int]string)

func main() {
	fmt.Println("🙌 Git pull ... ")

	// verif init command
	VerifInitCommand()

	// verif properties file
	path := VerifPropertiesFile()

	getRepositoriesGit(path)

	for _, repo := range repos {
		LaunchGitOperations(repo)
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
		if dir.IsDir() && isRealDir(dir.Name()) {
			//fmt.Println("          recur ", root, dir.Name())
			getRepositoriesGit(root + "/" + dir.Name())
		}
	}

}


func ReadDirectory(root string) ([]string, error) {
	var files []string
	fmt.Println("🔎 Git-pull is reading this directory ", root)

	directories, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println("error reading this directory ", root, " - err ", err)
		return files, err
	}

	for _, dir := range directories {
		if dir.IsDir() {

			if IsValidDir(dir.Name()) {
				// fmt.Println("\t is git dir ", dir.Name())
				LaunchGitOperations(root + dir.Name())
				
			} else {
				fmt.Println("======> ", dir.Name())
				//newDir := dir.Name()
				//fmt.Println("DIR newDir (root : ", root ,")", newDir)
				ReadDirectory(root+"/"+dir.Name())
			}
		}
	}

	return files, nil
}
