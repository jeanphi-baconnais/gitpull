package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var namePropertiesFile = ".gitpull.properties"

func VerifPropertiesFile() string {
	dat, err := ioutil.ReadFile(getPathPropertyFile())
	if err !=  nil {
		fmt.Println("Oh, it's the first time you lunch 'Git pull'. You need to take 2 seconds to configure your repo file.")
		return InitTool()
	}

	return string(dat)
}

func SavePropertiesFile(text string) {
	
	d1 := []byte(text)
	err := ioutil.WriteFile(getPathPropertyFile(), d1, 0644)
	if err !=  nil {
		panic(err)
	}
}

func getPathPropertyFile() string {
	home, _ := os.UserHomeDir()

	return home + "/" + namePropertiesFile
}