package utils

import (
	"bufio"
	"fmt"
	"jpbaconnais/git-pull/props"
	"os"
	"strings"
)

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
			props.SavePropertiesFile(text)
			return text
		}

	}
}