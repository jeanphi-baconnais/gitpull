package props

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func InitTool() string {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("🕸  Init git-pull tool")
		fmt.Print("Which directory contains git repositories ? ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if len(text) > 0 {
			fmt.Println("📝 Ok, lets go to add \"", text, "\" in the properties file. ")
			SavePropertiesFile(text)
			return text
		}

	}
}