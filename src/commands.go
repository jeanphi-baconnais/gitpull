package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func VerifInitCommand() {
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == commandInit {
			InitTool()
		} else if args[0] == commandHelp {
			PrintHelp()
		}
	}
}

func PrintHelp() {
	fmt.Println("---- Oh, you need some help. Let's me your introduce the Git pull tool 😃")
	fmt.Println("---- Git pull allows to make a simple git pull on all yours repositories save in one directory.")
	fmt.Println("---- To configure your repository directory, you can take gitpull help")
	fmt.Println("---- And next, you can only execute gitpull to pull all yours git repositories")
	fmt.Println("---- Enjoy ! 😎")
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